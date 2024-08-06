package db

import (
	"database/sql"

	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	var err error
	connStr := "postgresql://OmMahen:Ik7g1UWtOYlC@ep-sweet-glitter-836224.ap-southeast-1.aws.neon.tech/lostandfound_db?sslmode=require"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}
}
