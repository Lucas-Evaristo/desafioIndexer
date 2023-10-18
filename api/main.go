package main

import (
	"api-zincsearch-desafio/routes"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {
	var port string
	if len(os.Args) > 2 {
		port = ":" + os.Args[1]
	} else {
		port = ":4090"
	}
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Get("/search", routes.GetEmails)

	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal(err)
	}
}
