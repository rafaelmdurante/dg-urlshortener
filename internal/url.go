package internal

import (
	"errors"
	"github.com/rafaelmdurante/devgym-urlshortener/internal/base62"
	"net/url"
	"time"
)

type URL struct {
	ID        int       `json:"id"`
	TargetURL string    `json:"target_url"`
	URLCode   string    `json:"url_code"`
	CreatedAt time.Time `json:"created_at"`
}

var (
	ErrTargetURLInvalid = errors.New("target url is not valid")
)

func (u *URL) IsURLValid() (bool, error) {
	_, err := url.ParseRequestURI(u.TargetURL)

	if err != nil {
		return false, ErrTargetURLInvalid
	}

	return true, nil
}

func (u *URL) EncodeURL() URL {
	u.URLCode = base62.StdEncoding.Encode(u.ID)

	return *u
}

func (u *URL) DecodeURL(code string) URL {
	u.ID = base62.StdEncoding.Decode(code)

	return *u
}
