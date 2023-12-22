package internal

import (
	"errors"
	"github.com/rafaelmdurante/devgym-urlshortener/internal/base62"
	"net/url"
	"time"
)

type ShortenedURL struct {
	ID        int       `json:"id"`
	TargetURL string    `json:"target_url"`
	URLCode   string    `json:"url_code"`
	CreatedAt time.Time `json:"created_at"`
}

var ErrTargetURLInvalid = errors.New("target url is not valid")

func (u *ShortenedURL) IsURLValid() (bool, error) {
	_, err := url.ParseRequestURI(u.TargetURL)

	if err != nil {
		return false, ErrTargetURLInvalid
	}

	return true, nil
}

func (u *ShortenedURL) EncodeURL(id int) ShortenedURL {
	u.URLCode = base62.Encode(id)

	return *u
}
