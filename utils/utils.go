package utils

import (
	"os"
)

func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func PString(str string) *string {
	return &str
}

func PInt(i int) *int {
	return &i
}
