package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var store = map[string]string{}

type Shorten struct {
	URL string `json:"url"`
}

func shorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req Shorten
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, "missing URL", http.StatusBadRequest)
		return
	}

	url := req.URL
	if url == "" {
		http.Error(w, "missing URL", http.StatusBadRequest)
		return
	}

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}
	code := fmt.Sprintf("%d", len(store)+1)
	store[code] = url

	fmt.Fprintf(w, "short: http://localhost:8080/%s\n", code)
}

func redirect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not allowed", http.StatusBadRequest)
		return
	}
	code := r.URL.Path[1:]

	if url, ok := store[code]; ok {
		http.Redirect(w, r, url, http.StatusFound)
		return
	}

	http.NotFound(w, r)
}

func RunServer() {

	http.HandleFunc("/shorten", shorten)
	http.HandleFunc("/", redirect)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}
