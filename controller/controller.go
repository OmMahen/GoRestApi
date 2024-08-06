package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/OmMahen/GoRestApi/db"
	"github.com/OmMahen/GoRestApi/models"
	"github.com/gorilla/mux"
)

func AddMovie(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var movie models.Movie
	err := decoder.Decode(&movie)

	if err != nil {
		log.Println(err.Error())
	}

	query := `INSERT INTO movies (id, title, category, year, imdb_rating) VALUES ($1, $2, $3, $4, $5)`
	db.DB.Exec(query, movie.Id, movie.Title, movie.Category, movie.Year, movie.ImdbRating)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	rows, _ := db.DB.Query("SELECT id, title, category, year, imdb_rating FROM movies")

	var movies []models.Movie
	for rows.Next() {
		var movie models.Movie
		rows.Scan(&movie.Id, &movie.Title, &movie.Category, &movie.Year, &movie.ImdbRating)
		movies = append(movies, movie)
	}

	json.NewEncoder(w).Encode(movies)
}

func GetMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	id := mux.Vars(r)["id"]

	var movie models.Movie
	err := db.DB.QueryRow("SELECT * FROM movies WHERE id=$1", id).Scan(
		&movie.Id,
		&movie.Title,
		&movie.Category,
		&movie.Year,
		&movie.ImdbRating,
	)

	if err != nil {
		message := models.Message{Message: "Movie Not found"}
		json.NewEncoder(w).Encode(message)
	}

	json.NewEncoder(w).Encode(movie)
}

func DeleteMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	id := mux.Vars(r)["id"]
	result, err := db.DB.Exec("DELETE FROM movies WHERE id=$1", id)
	if err != nil {
		message := models.Message{Message: "Error deleting movie"}
		json.NewEncoder(w).Encode(message)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		message := models.Message{Message: "Error retrieving delete result"}
		json.NewEncoder(w).Encode(message)
	}
	if rowsAffected == 0 {
		message := models.Message{Message: "Movie not found"}
		json.NewEncoder(w).Encode(message)
	} else {
		message := models.Message{Message: "Movie successfully deleted"}
		json.NewEncoder(w).Encode(message)
	}
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var movie models.Movie

	json.NewDecoder(r.Body).Decode(&movie)

	query := `
		UPDATE movies
		SET title = $1, category = $2, year = $3, imdb_rating = $4
		WHERE id = $5
	`

	result, err := db.DB.Exec(query, movie.Title, movie.Category, movie.Year, movie.ImdbRating, movie.Id)

	if err != nil {
		message := models.Message{Message: "Error updating movie"}
		json.NewEncoder(w).Encode(message)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		message := models.Message{Message: "Error retrieving update result"}
		json.NewEncoder(w).Encode(message)
	}
	if rowsAffected == 0 {
		message := models.Message{Message: "Movie not found"}
		json.NewEncoder(w).Encode(message)
	} else {
		message := models.Message{Message: "Movie updated successfully"}
		json.NewEncoder(w).Encode(message)
	}
}
