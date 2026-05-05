package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	"url-shortner/db"
	"url-shortner/method"
	rds "url-shortner/redis"

	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
)

type Shorten struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
	Code     string `json:"code"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func shorten(w http.ResponseWriter, r *http.Request) {
	var req Shorten
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{
			Error: "Invalid JSON Body",
		})
		return
	}

	url := req.URL
	if url == "" {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{
			Error: "missing URL",
		})
		// http.Error(w, "missing URL", http.StatusBadRequest)
		return
	}

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}

	now := time.Now().UTC()
	var expiresAt time.Time
	var createdAt = now

	expiry := r.URL.Query().Get("expiry")
	if expiry != "" {
		minutes, err := strconv.Atoi(expiry)
		if err != nil {
			WriteJSON(w, http.StatusBadRequest, ErrorResponse{
				Error: "Invalid Expiry time",
			})
			return
		}

		t := now.Add(time.Duration(minutes) * time.Minute)
		expiresAt = t
	} else {
		t := now.Add(1 * time.Minute)
		expiresAt = t
	}
	var id int64

	err = db.Pool.QueryRow(context.Background(), "SELECT nextval(pg_get_serial_sequence('urls','id'))").Scan(&id)
	if err != nil {
		fmt.Println("DB error", err)
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{
			Error: "Failed to Generate ID",
		})
		// http.Error(w, "Failed to Generate ID", http.StatusInternalServerError)
		return
	}

	code := method.EncodeBase62(id)
	_, err = db.Pool.Exec(context.Background(), "INSERT INTO urls (id, short_code, long_url, expires_at, created_at) VALUES ($1, $2, $3, $4, $5)",
		id, code, url, expiresAt, createdAt,
	)

	if err != nil {
		fmt.Println("Insert Error:", err)
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{
			Error: "Insert Error",
		})
		// http.Error(w, "Insert Error", http.StatusInternalServerError)
		return
	}

	// fmt.Fprintf(w, "short: http://localhost:8080/%s\n", code)
	shortURL := fmt.Sprintf("http://localhost:8080/%s", code)
	w.Header().Set("Location", shortURL)
	WriteJSON(w, http.StatusCreated, ShortenResponse{
		ShortURL: shortURL,
		Code:     code,
	})

}

func redirect(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")

	// Redis
	ctx := r.Context()

	now := time.Now().UTC()

	if val, err := rds.Get(ctx, code); err == nil {
		fmt.Println("Redis Hit")
		http.Redirect(w, r, val, http.StatusFound)
		return
	} else if err != redis.Nil {
		fmt.Println("Redis Error:", err)
	}

	row := db.Pool.QueryRow(ctx,
		"SELECT long_url, expires_at FROM urls where short_code=$1", code,
	)
	var longURL string
	var expiresAt time.Time

	err := row.Scan(&longURL, &expiresAt)

	if err != nil {
		WriteJSON(w, http.StatusNotFound, ErrorResponse{
			Error: "URL Not Found",
		})
		// http.NotFound(w, r)
		return
	}

	if now.After(expiresAt) {
		WriteJSON(w, http.StatusGone, ErrorResponse{
			Error: "Link has been Expired",
		})
		// http.Error(w, "Link has been Expired", http.StatusGone)
		return
	}

	zeroTime := time.Time{}

	if expiresAt != zeroTime {
		ttl := expiresAt.Sub(now)
		fmt.Println("Time to Live is:", ttl)

		_ = rds.Set(ctx, code, longURL, ttl)
	} else {
		_ = rds.Set(ctx, code, longURL, 0)
	}
	http.Redirect(w, r, longURL, http.StatusFound)
}

func RegisterRoutes(r chi.Router) {
	r.Post("/shorten", shorten)
	r.Get("/{code}", redirect)
}
