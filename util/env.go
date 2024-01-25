package util

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	loadEnv()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured .env file. Err: %s", err)
	}
}

func Env(variable string) string {
	variableValue := os.Getenv(variable)
	return variableValue
}
