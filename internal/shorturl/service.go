package shorturl

import (
	"context"
	"errors"
	"github.com/rafaelmdurante/devgym-urlshortener/internal"
	"github.com/rafaelmdurante/devgym-urlshortener/internal/base62"
)

var (
	ErrTargetURLEmpty        = errors.New("target url cannot be empty")
	ErrShortURLAlreadyExists = errors.New("shortened url already exists for id")
	ErrInvalidURLCode        = errors.New("invalid url code")
	ErrCodeIsEmpty           = errors.New("code is empty")
	ErrURLNotFound           = errors.New("url not found")
)

type Service struct {
	Repository Repository
}

func (s Service) Create(ctx context.Context, u internal.ShortenedURL) (internal.ShortenedURL, error) {
	if u.TargetURL == "" {
		return internal.ShortenedURL{}, ErrTargetURLEmpty
	}

	if u.URLCode != "" {
		return internal.ShortenedURL{}, ErrShortURLAlreadyExists
	}

	// checks if uri is valid
	if valid, err := u.IsURLValid(); !valid && err != nil {
		return internal.ShortenedURL{}, err
	}

	// creates a row with empty url_code, ideally this step would be 'get valid id' from an id service, for instance
	// alternatively, if postgres could provide the next id number and reserve it, that would be awesome too
	r, err := s.Repository.Insert(ctx, u)
	if err != nil {
		return internal.ShortenedURL{}, err
	}

	// updates the empty url_code from step above
	return s.Repository.UpdateURLCode(ctx, r.EncodeURL(r.ID))
}

func (s Service) FindOneByCode(ctx context.Context, code string) (internal.ShortenedURL, error) {
	if code == "" {
		return internal.ShortenedURL{}, ErrCodeIsEmpty
	}

	id := base62.Decode(code)

	u, err := s.Repository.FindOneByID(ctx, id)

	if err != nil {
		return internal.ShortenedURL{}, err
	}

	return u, nil
}
