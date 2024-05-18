package db

import (
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

var DB *sqlx.DB

func InitDb() {
	// PG_URI should like this "postgres://username:pass@localhost:5432/db_name"
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connConfig, err := pgx.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Error connecting to the database", err)
		return
	}

	pgxdb := stdlib.OpenDB(*connConfig)
	DB = sqlx.NewDb(pgxdb, "pgx")
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(4)
	DB.SetConnMaxLifetime(time.Duration(30) * time.Minute)

	createTables()
	log.Println("eventBrite Database Connected")
}

func createTables() {
	createEventsTable := `
    CREATE TABLE IF NOT EXISTS events (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        description TEXT NOT NULL,
        location TEXT NOT NULL,
        dateTime TIMESTAMP NOT NULL,
        userId INTEGER
    )
    `
	_, err := DB.Exec(createEventsTable)
	if err != nil {
		log.Fatal(err)
		return
	}
}
