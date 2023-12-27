package shorturl

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rafaelmdurante/devgym-urlshortener/internal"
)

type Repository interface {
	Insert(ctx context.Context, shortURL internal.ShortenedURL) (internal.ShortenedURL, error)
	UpdateURLCode(ctx context.Context, shortURL internal.ShortenedURL) (internal.ShortenedURL, error)
	FindOneByID(ctx context.Context, id int) (internal.ShortenedURL, error)
}

type RepositoryPostgres struct {
	Conn *pgxpool.Pool
}

func (r *RepositoryPostgres) Insert(ctx context.Context, u internal.ShortenedURL) (internal.ShortenedURL, error) {
	err := r.Conn.QueryRow(
		ctx,
		"INSERT INTO url (target_url) VALUES ($1) RETURNING id, target_url, url_code, created_at",
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
		"UPDATE url SET url_code = $1 WHERE id = $2 RETURNING id, target_url, url_code, created_at",
		u.URLCode,
		u.ID,
	).Scan(&u.ID, &u.TargetURL, &u.URLCode, &u.CreatedAt)

	if err != nil {
		return internal.ShortenedURL{}, nil
	}

	return u, nil
}

func (r *RepositoryPostgres) FindOneByID(ctx context.Context, id int) (internal.ShortenedURL, error) {
	u := internal.ShortenedURL{}
	err := r.Conn.QueryRow(
		ctx,
		"SELECT id, target_url FROM url WHERE id = $1",
		id,
	).Scan(&u.ID, &u.TargetURL)

	if errors.Is(err, pgx.ErrNoRows) {
		return internal.ShortenedURL{}, ErrURLNotFound
	}

	if err != nil {
		return internal.ShortenedURL{}, err
	}

	return u, nil
}
