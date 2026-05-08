package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type VisitedStore struct {
	client *redis.Client
	ctx    context.Context
}

func NewVisitedStore(client *redis.Client) *VisitedStore {
	return &VisitedStore{
		client: client,
		ctx:    context.Background(),
	}
}

const visitedKey = "visited_urls"
const inProgressKey = "in_progress_urls"

func (r *VisitedStore) Seen(url string) bool {
	added, err := r.client.SAdd(
		r.ctx, visitedKey, url,
	).Result()

	if err != nil {
		return false
	}
	return added == 0
}

func (r *VisitedStore) MarkInProgress(url string) bool {
	added, err := r.client.SAdd(
		r.ctx, inProgressKey, url,
	).Result()

	if err != nil {
		return false
	}
	return added > 0
}

func (r *VisitedStore) RemoveInProgress(url string) error {
	return r.client.SRem(r.ctx, inProgressKey, url).Err()
}

func (r *VisitedStore) Cleanup() error {
	return r.client.Del(r.ctx, visitedKey).Err()
}
