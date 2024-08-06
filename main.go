package main

import (
	"log"
	"net/http"
	"os"

	"github.com/OmMahen/GoRestApi/db"
	"github.com/OmMahen/GoRestApi/routes"
	"github.com/joho/godotenv"
)

func main() {
	db.Init()
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	router := routes.MovieRoutes()

	http.Handle("/api", router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not set
	}

	log.Println("Listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
