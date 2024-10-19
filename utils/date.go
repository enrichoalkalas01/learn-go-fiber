package utils

import (
	"time"
)

type expiredType string

const (
	Seconds expiredType = "seconds"
	Minutes expiredType = "minutes"
	Hours   expiredType = "hours"
	Days    expiredType = "days"
	Months  expiredType = "months"
	Years   expiredType = "years"
)

type futureDateParams struct {
	ExpiredType expiredType
	Value       int
}

type resultGetFutureDate struct {
	FormattedDate *string
	Milliseconds  *int64
	FutureTime    *time.Time
}

// func getFutureDate(params futureDateParams) (string, int64, string) {
func getFutureDate(params futureDateParams) resultGetFutureDate {
	currentTime := time.Now()
	var futureTime time.Time

	switch params.ExpiredType {
	case Seconds:
		futureTime = currentTime.Add(time.Duration(params.Value) * time.Second)

	case Minutes:
		futureTime = currentTime.Add(time.Duration(params.Value) * time.Minute)

	case Hours:
		futureTime = currentTime.Add(time.Duration(params.Value) * time.Hour)

	case Days:
		futureTime = currentTime.AddDate(0, 0, params.Value)

	case Months:
		futureTime = currentTime.AddDate(0, params.Value, 0)

	case Years:
		futureTime = currentTime.AddDate(params.Value, 0, 0)

	default:
		futureTime = currentTime
	}

	formattedDate := time.UnixMilli(futureTime.UnixMilli()).Format("2006-01-02 15:04:05")
	milliseconds := futureTime.UnixMilli()

	return resultGetFutureDate{
		FormattedDate: &formattedDate,
		Milliseconds:  &milliseconds,
		FutureTime:    &futureTime,
	}
}
