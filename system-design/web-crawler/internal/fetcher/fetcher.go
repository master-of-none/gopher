package fetcher

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: 5 * time.Second,
}

func Fetch(ctx context.Context, url string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("Non 200 response")
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}
