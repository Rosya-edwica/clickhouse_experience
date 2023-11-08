package tools

import (
	"strings"
	"time"
)

const DefaultDate = "2001-01-01 00:00:00.000"

func ConvertDataToDateTime(date string) string {
	if strings.Contains(date, ":") {
		return DefaultDate
	}
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return DefaultDate
	}
	t = t.Add(time.Duration(time.Now().Hour()) + time.Duration(time.Now().Minute()) + time.Duration(time.Now().Second()))
	newDate := t.Format("2006-01-02 15:04:05")
	return newDate
}
