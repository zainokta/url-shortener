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

func GetShorten(ctx context.Context, shortenURL string) (Shortener, error) {
	query := "SELECT uuid, original_url, shorten_url FROM shorteners WHERE shorten_url=@shorten_url"
	args := pgx.NamedArgs{
		"shorten_url": shortenURL,
	}

	rows, err := database.DB.Query(ctx, query, args)
	if err != nil {
		log.Println(err.Error())
		return Shortener{}, err
	}

	shortener, err := pgx.CollectOneRow(rows, func(row pgx.CollectableRow) (Shortener, error) {
		var s Shortener
		err := row.Scan(&s.UUID, &s.OriginalURL, &s.ShortenURL)
		if err != nil {
			log.Println(err.Error())
			return s, err
		}
		return s, nil
	})
	if err != nil {
		if err == pgx.ErrNoRows {
			return Shortener{}, pgx.ErrNoRows
		}

		log.Println(err.Error())
		return Shortener{}, err
	}

	return shortener, nil
}

func StoreShortenerLog(ctx context.Context, shortenerId string, newUrl string, oldUrl string) error {
	query := "INSERT INTO shortener_logs (shortener_id, new_url, old_url) VALUES (@shortener_id, @new_url, @old_url)"
	args := pgx.NamedArgs{
		"shortener_id": shortenerId,
		"new_url":      newUrl,
		"old_url":      oldUrl,
	}

	_, err := database.DB.Exec(ctx, query, args)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func UpdateShortener(ctx context.Context, newUrl string, shortenUrl string) error {
	query := "UPDATE shorteners SET original_url=@new_url WHERE shorten_url=@shorten_url"
	args := pgx.NamedArgs{
		"new_url":     newUrl,
		"shorten_url": shortenUrl,
	}

	_, err := database.DB.Exec(ctx, query, args)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
