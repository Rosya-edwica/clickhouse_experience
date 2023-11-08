package entities

type Vacancy struct {
	Id          int     `db:"id"`
	VacancyId   string  `db:"vacancy_id"`
	Url         string  `db:"url"`
	Name        string  `db:"name"`
	CityId      int     `db:"city_id"`
	PositionId  int     `db:"position_id"`
	Areas       string  `db:"prof_areas"`
	Specs       string  `db:"specs"`
	Experience  string  `db:"experience"`
	SalaryFrom  float64 `db:"salary_from"`
	SalaryTo    float64 `db:"salary_to"`
	Skills      string  `db:"key_skills"`
	VacancyDate string  `db:"vacancy_date"`
	ParsingDate string  `db:"parsing_date"`
	Platform    string  `db:"platform"`
}
