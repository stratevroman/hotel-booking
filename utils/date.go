package utils

import (
	"errors"
	"time"
)

var ErrIncorrectDates = errors.New("dates is incorrect")

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func DaysBetween(from time.Time, to time.Time) ([]time.Time, error) {
	if from.After(to) {
		return nil, ErrIncorrectDates
	}

	days := make([]time.Time, 0)
	for d := toDay(from); !d.After(toDay(to)); d = d.AddDate(0, 0, 1) {
		days = append(days, d)
	}

	return days, nil
}

func toDay(timestamp time.Time) time.Time {
	return time.Date(timestamp.Year(), timestamp.Month(), timestamp.Day(), 0, 0, 0, 0, time.UTC)
}
