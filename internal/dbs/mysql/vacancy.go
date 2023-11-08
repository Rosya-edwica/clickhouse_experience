package mysql

import (
	"fmt"
	"github.com/Rosya-edwica/clickhouse_experience/internal/entities"
	"github.com/Rosya-edwica/clickhouse_experience/internal/models"

	"github.com/go-faster/errors"
	"github.com/jmoiron/sqlx"
)

func GetVacancies(conn *sqlx.DB, lastId int) ([]models.Vacancy, error) {
	rawVacancies := make([]entities.Vacancy, 0)

	err := conn.Select(&rawVacancies, fmt.Sprintf(`
		SELECT id, vacancy_id, url, name, city_id, position_id, prof_areas, specs, experience, salary_from, salary_to, 
		key_skills, vacancy_date, parsing_date, platform 
		FROM h_vacancy WHERE id > %d 
		ORDER BY id
		LIMIT 10000;`,
		lastId),
	)
	if err != nil {
		return nil, errors.Wrap(err, "clickhouse-select vacancy failed")
	}
	return models.NewVacancies(rawVacancies), nil
}
