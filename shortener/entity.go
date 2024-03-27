package shortener

import (
	"errors"
	"net/url"
	"time"
)

type Shortener struct {
	UUID        string    `json:"uuid" db:"uuid"`
	OriginalURL string    `json:"originalUrl" db:"original_url"`
	ShortenURL  string    `json:"shortenUrl" db:"shorten_url"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}

var (
	ErrInvalidLength = errors.New("url length must not be more than 256")
	ErrNotValidURL   = errors.New("requested url is not valid url")
)

func (s Shortener) Validate() error {
	if len(s.OriginalURL) > 256 {
		return ErrInvalidLength
	}

	if len(s.ShortenURL) > 256 {
		return ErrInvalidLength
	}

	original, err := url.Parse(s.OriginalURL)
	if err != nil && original.Scheme == "" && original.Host == "" {
		return ErrNotValidURL
	}

	return nil
}

type ShortenerLog struct {
	ShortenerID string `json:"shortenerId" db:"shortener_id"`
	NewURL      string `json:"newUrl" db:"new_url"`
	OldURL      string `json:"oldUrl" db:"old_url"`
}
