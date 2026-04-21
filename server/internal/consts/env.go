package consts

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var DB_PORT string
var DB_HOST string
var DB_EXT string
var DB_NAME string
var DB_USER string
var DB_PWD string

var REDIS_PORT string
var REDIS_HOST string
var REDIS_EXT string
var REDIS_PWD string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB_PORT = os.Getenv("DB_PORT")
	DB_HOST = os.Getenv("DB_HOST")
	DB_EXT = os.Getenv("DB_EXT")
	DB_NAME = os.Getenv("DB_NAME")
	DB_USER = os.Getenv("DB_USER")
	DB_PWD = os.Getenv("DB_PWD")

	REDIS_PORT = os.Getenv("REDIS_PORT")
	REDIS_HOST = os.Getenv("REDIS_HOST")
	REDIS_EXT = os.Getenv("REDIS_EXT")
	REDIS_PWD = os.Getenv("REDIS_PWD")
}
