package db

import (
	"fmt"
	"log"
	"user-service/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgres(cfg *config.Config) *sqlx.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.Name, cfg.DB.SSLMode)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	log.Println("Connected to PostgreSQL")
	return db
}
