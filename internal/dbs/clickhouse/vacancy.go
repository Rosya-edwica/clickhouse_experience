package clickhouse

import (
	"context"
	"github.com/Rosya-edwica/clickhouse_experience/internal/models"
	"github.com/Rosya-edwica/clickhouse_experience/tools"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/pkg/errors"

	"strings"
)

func SaveVacancies(conn driver.Conn, items []models.Vacancy) error {
	batch, err := conn.PrepareBatch(context.Background(), "INSERT INTO vacancy")
	if err != nil {
		return errors.Wrap(err, "clickhouse-creating batch failed")
	}

	for _, i := range items {
		err := batch.Append(
			i.Id,
			i.VacancyId,
			i.Platform,
			i.Name,
			i.PositionId,
			i.CityId,
			i.Experience,
			i.SalaryFrom,
			i.SalaryTo,
			tools.ConvertDataToDateTime(i.VacancyDate),
			tools.ConvertDataToDateTime(i.ParsingDate),
			strings.Split(i.Skills, "|"),
			strings.Split(i.Areas, "|"),
			strings.Split(i.Specs, "|"),
		)
		if err != nil {
			return errors.Wrap(err, "clickhouse-adding to batch failed")
		}
	}
	err = batch.Send()
	if err != nil {
		return errors.Wrap(err, "clickhouse-sending batch failed")
	}
	return nil
}
