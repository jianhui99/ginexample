package config

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var once sync.Once

func Init() {
	once.Do(func() {
		godotenv.Load()
	})
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
