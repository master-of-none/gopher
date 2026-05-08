package frontier

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

const queueKey = "crawler:queue"

type RedisQueue struct {
	client *redis.Client
}

func NewRedisQueue(client *redis.Client) *RedisQueue {
	return &RedisQueue{client: client}
}

func (q *RedisQueue) Push(task CrawlTask) error {
	payload, err := json.Marshal(task)
	if err != nil {
		return err
	}
	return q.client.RPush(context.Background(), queueKey, payload).Err()
}

func (q *RedisQueue) Pop() CrawlTask {
	result, err := q.client.LPop(context.Background(), queueKey).Result()
	if err != nil {
		return CrawlTask{}
	}
	var task CrawlTask
	if err := json.Unmarshal([]byte(result), &task); err != nil {
		return CrawlTask{}
	}
	return task
}

func (q *RedisQueue) PopContext(ctx context.Context) (CrawlTask, bool) {
	result, err := q.client.BLPop(ctx, 0, queueKey).Result()
	if err != nil {
		return CrawlTask{}, false
	}
	var task CrawlTask
	if err := json.Unmarshal([]byte(result[1]), &task); err != nil {
		return CrawlTask{}, false
	}
	return task, true
}

func (q *RedisQueue) Cleanup() error {
	return q.client.Del(context.Background(), queueKey).Err()
}
