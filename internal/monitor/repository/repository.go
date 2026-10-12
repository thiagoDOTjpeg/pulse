package repository

import (
	"context"
	"fmt"
	"log/slog"
	"thiagodotjpeg/pulse/internal/monitor/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UrlRepository struct {
	pool *pgxpool.Pool
	logger *slog.Logger
}

func newUrlRepository(pool *pgxpool.Pool, logger *slog.Logger) *UrlRepository {
	return &UrlRepository{pool: pool, logger: logger }
}

func (r *UrlRepository) Create(ctx context.Context, url *model.Url) (*model.Url, error) {
	const query = `
		INSERT INTO url (url, created_at, updated_at, next_verification_at)
		VALUES ($1, $2, $3, $4)
		RETURN *;
	`

	_, err := r.pool.Query(ctx, query, url.Url, url.CreatedAt, url.UpdatedAt, url.NextVerificationAt)
	if err != nil {
		r.logger.Error("Ocorreu um erro ao criar o monitoramento", "error", err)
		return nil, fmt.Errorf("Ocorreu um erro ao criar o monitoramento: %w", err)
	}

	return url, nil
}