package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type URLRepository struct {
	pool *pgxpool.Pool
	ctx  context.Context
}

func NewURLRepository(pool *pgxpool.Pool) *URLRepository {
	return &URLRepository{
		pool: pool,
		ctx:  context.Background(),
	}
}

func (r *URLRepository) InsertURL(url string, depth int) error {
	_, err := r.pool.Exec(
		r.ctx, `INSERT INTO urls(url, depth) VALUES($1,$2) ON CONFLICT (url) DO NOTHING`, url, depth,
	)
	return err
}

func (r *URLRepository) Exists(url string) (bool, error) {
	var exists bool

	err := r.pool.QueryRow(
		r.ctx, `SELECT EXISTS (SELECT 1 from urls WHERE url=$1)`, url,
	).Scan(&exists)
	return exists, err
}

func (r *URLRepository) MarkCrawled(url string) error {
	_, err := r.pool.Exec(
		r.ctx, `UPDATE urls SET status='success', crawled_at=NOW() WHERE url=$1`, url,
	)
	return err
}

func (r *URLRepository) Truncate() error {
	_, err := r.pool.Exec(r.ctx, `DELETE FROM urls`)
	return err
}
