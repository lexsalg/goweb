package config

import (
	"os"
	"strconv"
)

func StringEnv(env string, defEnv string) string {
	val := os.Getenv(env)
	if val == "" {
		return defEnv
	}
	return val
}

func IntEnv(env string, defEnv int) int {
	val := os.Getenv(env)
	if val == "" {
		return defEnv
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		return defEnv
	}
	return i
}

func BoolEnv(env string, defEnv bool) bool {
	val := os.Getenv(env)
	if val == "" {
		return defEnv
	}
	b, err := strconv.ParseBool(val)
	if err != nil {
		return defEnv
	}
	return b
}
