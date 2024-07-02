package storage

import (
	"database/sql"
	"fmt"

	"github.com/saalcazar/surveys/model"
)

type scanner interface {
	Scan(dest ...any) error
}

const (
	psqlCreateSurvey = `SELECT create_survey($1, $2, $3, $4, $5)`
	psqlGetAllSurvey = `SELECT * FROM surveys`
)

type PsqlSurvey struct {
	db *sql.DB
}

func NewPsqlSurvey(db *sql.DB) *PsqlSurvey {
	return &PsqlSurvey{db}
}

func (p *PsqlSurvey) Create(m *model.Survey) error {
	stmt, err := p.db.Prepare(psqlCreateSurvey)
	if err != nil {
		return err
	}

	defer stmt.Close()
	err = stmt.QueryRow(
		m.Mail,
		m.Place,
		m.Age,
		m.Gender,
		m.Option,
	).Scan(&m.Mail)
	if err != nil {
		return err
	}
	fmt.Println("Se crearon los datos de manera correcta")
	return nil
}

func (p *PsqlSurvey) GetAll() (model.Surveys, error) {
	stmt, err := p.db.Prepare(psqlGetAllSurvey)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	ms := make(model.Surveys, 0)
	for rows.Next() {
		p, err := scanRowSurveys(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, *p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ms, nil
}

func scanRowSurveys(s scanner) (*model.Survey, error) {
	m := &model.Survey{}
	err := s.Scan(
		&m.ID,
		&m.Date,
		&m.Mail,
		&m.Place,
		&m.Age,
		&m.Gender,
		&m.Option,
	)
	if err != nil {
		return &model.Survey{}, err
	}
	return m, nil
}
