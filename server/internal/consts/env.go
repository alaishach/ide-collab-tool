package consts

import (
	"log"
	"os"
	"slices"

	"github.com/joho/godotenv"
)

var ENV string

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

var SERVER_DOMAIN string

func validValue(varName string, validValues []string) string {
	value := os.Getenv(varName)
	if value == "" {
		panic("Env variable '" + varName + "' has not been defined in .env")
	}
	if !slices.Contains(validValues, value) {
		panic("Env variable '" + varName + "' has not been set to a valid value")
	}
	return value
}

func getEnv(varName string) string {
	value := os.Getenv(varName)
	if value == "" {
		panic("Env variable '" + varName + "' has not been defined in .env")
	}
	return value
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ENV = validValue("ENV", []string{"dev", "prod"})

	DB_PORT = getEnv("DB_PORT")
	println("PORT!!!!!!!!!!: ", DB_PORT)
	DB_HOST = getEnv("DB_HOST")
	DB_EXT = getEnv("DB_EXT")
	DB_NAME = getEnv("DB_NAME")
	DB_USER = getEnv("DB_USER")
	DB_PWD = getEnv("DB_PWD")

	REDIS_PORT = getEnv("REDIS_PORT")
	REDIS_HOST = getEnv("REDIS_HOST")
	REDIS_EXT = getEnv("REDIS_EXT")
	REDIS_PWD = getEnv("REDIS_PWD")

	SERVER_DOMAIN = getEnv("SERVER_DOMAIN")
}
