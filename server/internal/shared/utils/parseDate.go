package utils

import "time"

func ParseDate(dateStr string) (time.Time, error) {
	layout := "02/01/2006" // Формат даты day/month/year
	return time.Parse(layout, dateStr)
}