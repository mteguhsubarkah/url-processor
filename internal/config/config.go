package config
import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

func Load() {
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using system environment variables")
    }
}

func Get(key string) string {
    return os.Getenv(key)
}