package models

import (
	"github.com/Rosya-edwica/clickhouse_experience/internal/entities"
)

type Vacancy struct {
	Id          int
	VacancyId   string
	Url         string
	Name        string
	CityId      int
	PositionId  int
	Areas       string
	Specs       string
	Experience  string
	SalaryFrom  float64
	SalaryTo    float64
	Skills      string
	VacancyDate string
	ParsingDate string
	Platform    string
}

func NewVacancies(rawVacancies []entities.Vacancy) (vacancies []Vacancy) {
	for _, item := range rawVacancies {
		vacancies = append(vacancies, Vacancy{
			Id:          item.Id,
			VacancyId:   item.VacancyId,
			Url:         item.Url,
			Name:        item.Name,
			CityId:      item.CityId,
			PositionId:  item.PositionId,
			Areas:       item.Areas,
			Specs:       item.Specs,
			Experience:  item.Experience,
			SalaryFrom:  item.SalaryFrom,
			SalaryTo:    item.SalaryTo,
			Skills:      item.Skills,
			VacancyDate: item.VacancyDate,
			ParsingDate: item.ParsingDate,
			Platform:    item.Platform,
		})
	}
	return
}
