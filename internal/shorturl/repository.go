package shorturl

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rafaelmdurante/devgym-urlshortener/internal"
)

type Repository interface {
	Insert(ctx context.Context, shortURL internal.ShortenedURL) (internal.ShortenedURL, error)
	UpdateURLCode(ctx context.Context, shortURL internal.ShortenedURL) (internal.ShortenedURL, error)
}

type RepositoryPostgres struct {
	Conn *pgxpool.Pool
}

func (r *RepositoryPostgres) Insert(ctx context.Context, u internal.ShortenedURL) (internal.ShortenedURL, error) {
	err := r.Conn.QueryRow(
		ctx,
		"INSERT INTO short_url (target_url) VALUES ($1) RETURNING id, target_url, url_code, created_at",
		u.TargetURL,
	).Scan(&u.ID, &u.TargetURL, &u.URLCode, &u.CreatedAt)

	if err != nil {
		return internal.ShortenedURL{}, err
	}

	return u, nil
}

func (r *RepositoryPostgres) UpdateURLCode(ctx context.Context, u internal.ShortenedURL) (internal.ShortenedURL, error) {
	err := r.Conn.QueryRow(
		ctx,
		"UPDATE short_url SET url_code = $1 WHERE id = $2 RETURNING id, target_url, url_code, created_at",
		u.URLCode,
		u.ID,
	).Scan(&u.ID, &u.TargetURL, &u.URLCode, &u.CreatedAt)

	if err != nil {
		return internal.ShortenedURL{}, nil
	}

	return u, nil
}
