package mw

import (
	"net"
	"net/http"
	"time"
	"url-shortner/handler"
	rds "url-shortner/redis"
)

const (
	limit  = 10
	window = 60 * time.Second
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func RateLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		key := "rate_limit:" + ip

		count, err := rds.Client.Incr(ctx, key).Result()

		if err != nil {
			handler.WriteJSON(w, http.StatusInternalServerError, ErrorResponse{
				Error: "Internal Server Error",
			})
			return
		}

		if count == 1 {
			rds.Client.Expire(ctx, key, window)
		}

		if count > limit {
			w.Header().Set("Retry-After", "60")
			handler.WriteJSON(w, http.StatusTooManyRequests, ErrorResponse{
				Error: "Too many Requests",
			})
			return
		}
		next.ServeHTTP(w, r)
	})
}
