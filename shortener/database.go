package shortener

import (
	"context"
	"log"
	"url-shortener/database"

	"github.com/jackc/pgx/v5"
)

func StoreShortener(ctx context.Context, req Shortener) error {
	query := "INSERT INTO shorteners (original_url, shorten_url) VALUES (@original_url, @shorten_url)"
	args := pgx.NamedArgs{
		"original_url": req.OriginalURL,
		"shorten_url":  req.ShortenURL,
	}

	_, err := database.DB.Exec(ctx, query, args)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
