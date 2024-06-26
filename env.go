package goutils

import "os"

func GetEnv(key string, def string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return def
}

func GetEnvOrPanic(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	panic("Environment variable " + key + " not found")
}
