package config

import (
	"os"
	"strconv"
	"strings"
)

func getString(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		n, _ := strconv.Atoi(value)
		return n
	}

	return fallback
}

func getStringList(key string, fallback []string) []string {
	if value, ok := os.LookupEnv(key); ok {
		return strings.Split(value, ",")
	}

	return fallback
}
