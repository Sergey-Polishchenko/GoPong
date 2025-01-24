package shared

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var (
	loadOnce sync.Once
	loadErr  error
)

func loadEnv() error {
	loadOnce.Do(func() {
		loadErr = godotenv.Load()
		if loadErr != nil {
			loadErr = fmt.Errorf("cant load .env: %s", loadErr)
		}
	})
	return loadErr
}

func LoadVar(name string) (string, error) {
	if err := loadEnv(); err != nil {
		return "", err
	}

	value, contains := os.LookupEnv(name)

	if !contains {
		return "", fmt.Errorf("environment variable %s not found; check your .env file", name)
	}
	return value, nil
}
