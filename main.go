package main

import (
	"net/http"

	"github.com/saalcazar/surveys/handler"
	"github.com/saalcazar/surveys/storage"
)

func main() {
	storage.NewPostgresDB()

	storageSurvey := storage.NewPsqlSurvey(storage.Pool())
	serviceSurvey := handler.NewServiceSurvey(storageSurvey)

	mux := http.NewServeMux()

	handler.RouteSurvey(mux, serviceSurvey)

	http.ListenAndServe(":8080", mux)

}
