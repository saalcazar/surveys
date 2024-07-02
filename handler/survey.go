package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/saalcazar/surveys/model"
)

type ServiceSurvey struct {
	storageSurvey StorageSurvey
}

func NewServiceSurvey(s StorageSurvey) *ServiceSurvey {
	return &ServiceSurvey{s}
}

func (s *ServiceSurvey) Create(p *model.Survey) error {
	return s.storageSurvey.Create(p)
}

func (s *ServiceSurvey) GetAll() (model.Surveys, error) {
	return s.storageSurvey.GetAll()
}

func (s *ServiceSurvey) create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data := model.Survey{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		response := newResponse(Error, "Los datos no tienen la estructura correcta", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = s.Create(&data)
	if err != nil {
		log.Fatalf("Error: %+v", err)
		response := newResponse(Error, "Hubo un problema al subir los datos", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "Datos subidos al servidor", nil)
	responseJSON(w, http.StatusCreated, response)
}

func (s *ServiceSurvey) getAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response := newResponse(Error, "Método no permitido", nil)
		responseJSON(w, http.StatusBadRequest, response)
		return
	}

	data, err := s.GetAll()
	if err != nil {
		response := newResponse(Error, "Hubo un problema al obtener los datos", nil)
		responseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := newResponse(Message, "OK", data)
	responseJSON(w, http.StatusOK, response)

}
