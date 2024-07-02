package handler

import "github.com/saalcazar/surveys/model"

type StorageSurvey interface {
	Create(*model.Survey) error
	// GetAll() (model.Surveys, error)
}
