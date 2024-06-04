package database

import (
	"context"
	"fmt"
	"net/url"
	"url-shortener/env"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func New(ctx context.Context) error {
	pass := url.QueryEscape(env.GetEnv("DB_PASSWORD", "Master@123"))
	dsn := fmt.Sprintf("postgres://%s:%s@%s/%s",
		env.GetEnv("DB_USERNAME", "postgres"),
		pass,
		env.GetEnv("DB_HOST", "localhost:5432"),
		env.GetEnv("DB_NAME", "link_shortener"),
	)

	conn, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return err
	}

	DB = conn

	return nil
}
