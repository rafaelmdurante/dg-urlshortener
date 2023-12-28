package url

import (
	"context"
	"errors"
	"fmt"
	"github.com/rafaelmdurante/devgym-urlshortener/internal"
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

func (s Service) Create(ctx context.Context, input internal.URL) (internal.URL, error) {
	if input.TargetURL == "" {
		return internal.URL{}, ErrTargetURLEmpty
	}

	if input.URLCode != "" {
		return internal.URL{}, ErrShortURLAlreadyExists
	}

	// checks if uri is valid
	if valid, err := input.IsURLValid(); !valid && err != nil {
		return internal.URL{}, err
	}

	// creates a row with empty url_code, ideally this step would be 'get valid id' from an id service, for instance
	// alternatively, if postgres could provide the next id number and reserve it, that would be awesome too
	u, err := s.Repository.Insert(ctx, input)
	if err != nil {
		return internal.URL{}, err
	}

	fmt.Sprintln("going to create")

	// updates the empty url_code from step above
	return s.Repository.UpdateURLCode(ctx, u.EncodeURL())
}

func (s Service) FindOneByCode(ctx context.Context, code string) (internal.URL, error) {
	if code == "" {
		return internal.URL{}, ErrCodeIsEmpty
	}

	var u internal.URL
	var err error

	u, err = s.Repository.FindOneByID(ctx, u.DecodeURL(code).ID)

	if err != nil {
		return internal.URL{}, err
	}

	return u, nil
}
