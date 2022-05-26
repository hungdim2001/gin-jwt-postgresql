package database

import (
	"os"

	"github.com/joho/godotenv"
)

func config(envStr string) string {
	err := godotenv.Load(".env")
	if err != nil {
		panic("can't load env file")
	}
	return os.Getenv(envStr)

}
