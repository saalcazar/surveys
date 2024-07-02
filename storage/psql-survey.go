package storage

import (
	"database/sql"
	"fmt"

	"github.com/saalcazar/surveys/model"
)

// type scanner interface {
// 	Scan(dest ...any) error
// }

const (
	psqlCreateSurvey = `SELECT create_survey($1, $2, $3, $4, $5)`
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
