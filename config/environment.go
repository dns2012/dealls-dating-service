package config

import "os"

func ENV(key, fallback string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		return fallback
	}

	return value
}
