package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"url-shortner/cron"
	"url-shortner/db"
	"url-shortner/handler"
	"url-shortner/mw"
	rds "url-shortner/redis"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	db.InitDB()
	defer db.CloseDB()
	rds.InitRedis()
	defer rds.CloseRedis()
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(mw.RateLimiter)
	r.Use(mw.CORSMiddleware)
	r.Use(middleware.Recoverer)

	handler.RegisterRoutes(r)
	// mux := http.NewServeMux()
	// handler.RegisterRoutes(mux)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	ctx, cancel := context.WithCancel(context.Background())
	cron.StartCleanupWorker(ctx, 1*time.Minute)
	go func() {
		fmt.Println("Server starting on 8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	fmt.Println("Shutting down...")

	cancel()

	shutDownCtx, shutDownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutDownCancel()
	server.Shutdown(shutDownCtx)
	if err := db.DeleteExpired(shutDownCtx); err != nil {
		fmt.Println("Cleanup Error:", err)
	}

	fmt.Println("Cleanup complete")
}
