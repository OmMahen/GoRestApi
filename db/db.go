package db

import (
	"database/sql"

	"github.com/OmMahen/GoRestApi/models"

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

var MovieList = []models.Movie{

	{
		Id:         "1",
		Title:      "The Godfather",
		Year:       "1972",
		Category:   "Cime/Drama",
		ImdbRating: "9.2/10",
	},
	{
		Id:         "2",
		Title:      "The English Patient",
		Year:       "1996",
		Category:   "Romance/Drama",
		ImdbRating: "7.4/10",
	}, {
		Id:         "3",
		Title:      "The Greate Gatsby",
		Year:       "2013",
		Category:   "Romance/Drama",
		ImdbRating: "7.2/10",
	},
}
