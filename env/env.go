package env

import (
	"os"
	"strconv"
)

func GetEnv(key string, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	return fallback
}

func GetEnvInt(key string, fallback int) (int, error) {
	if val := os.Getenv(key); val != "" {
		valInt, err := strconv.Atoi(val)
		if err != nil {
			return 0, err
		}

		return valInt, nil
	}

	return fallback, nil
}
