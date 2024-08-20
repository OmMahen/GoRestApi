package db

import (
	"database/sql"

	"log"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"

	"os"
)

var DB *sql.DB

func Init() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	connStr := os.Getenv("DATABASE_CONNECTION_STRING")
	if connStr == "" {
		log.Fatalf("API_KEY environment variable is not set")
		return
	}

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}
}
