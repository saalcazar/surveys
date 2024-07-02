package handler

import "net/http"

func RouteSurvey(mux *http.ServeMux, storageSurvey StorageSurvey) {
	h := NewServiceSurvey(storageSurvey)
	mux.HandleFunc("/v1/survey/create", h.create)
	mux.HandleFunc("/v1/survey/getall", h.getAll)
}
