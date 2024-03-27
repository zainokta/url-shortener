package shortener

import (
	"errors"
	"time"
)

type Shortener struct {
	UUID        string    `json:"uuid"`
	OriginalURL string    `json:"originalUrl"`
	ShortenURL  string    `json:"shortenUrl"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

var (
	ErrInvalidLength = errors.New("url length must not be more than 256")
)

func (s Shortener) Validate() error {
	if len(s.OriginalURL) > 256 {
		return ErrInvalidLength
	}

	if len(s.ShortenURL) > 256 {
		return ErrInvalidLength
	}

	return nil
}

type ShortenerLog struct {
	ShortenerID string `json:"shortenerId"`
	NewURL      string `json:"newUrl"`
}
