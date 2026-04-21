// Package pg is dope
package pg

import (
	"fmt"
	"log"
	"server/internal/consts"
	"server/internal/utils/logger"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func init() {
	dsn := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", consts.DB_USER, consts.DB_NAME, consts.DB_PWD, consts.DB_HOST, consts.DB_PORT)

	var err error
	DB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to create database connection: %v", err)
	}

	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(10)
	if err = DB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	logger.Logger.Info("Database connected successfully")
}
