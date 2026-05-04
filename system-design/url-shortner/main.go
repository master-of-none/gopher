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

	"github.com/go-chi/chi/v5"
)

func main() {
	db.InitDB()

	r := chi.NewRouter()
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

	db.CloseDB()

	fmt.Println("Cleanup complete")
}
