package main

import (
	"net/http"

	"github.com/rs/cors"
	"github.com/saalcazar/surveys/handler"
	"github.com/saalcazar/surveys/storage"
)

func main() {
	storage.NewPostgresDB()

	storageSurvey := storage.NewPsqlSurvey(storage.Pool())
	serviceSurvey := handler.NewServiceSurvey(storageSurvey)

	mux := http.NewServeMux()

	handler.RouteSurvey(mux, serviceSurvey)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8080/v1/", "http://localhost:5173"},
		AllowedMethods: []string{"DELETE", "GET", "POST", "PUT"},
	})

	handler := c.Handler(mux)

	http.ListenAndServe(":8080", handler)

}
