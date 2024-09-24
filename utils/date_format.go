package utils

import (
	"fmt"
	"hkn-be/constants"
	"time"
)

func ValidateDateFormat(date *string, flag string) error {
	if date == nil || *date == "" {
		*date = time.Now().Format(time.DateOnly)
	}
	_, err := time.Parse(time.DateOnly, *date)
	if flag == constants.StartFlag {
		*date = fmt.Sprintf("%s 00:00:00", *date)
	}
	if flag == constants.EndFlag {
		*date = fmt.Sprintf("%s 23:59:59", *date)
	}
	return err
}
func ValidateDateTimeFormat(dateTime *string) error {
	if dateTime == nil || *dateTime == "" {
		*dateTime = time.Now().Format(time.DateTime)
	}
	_, err := time.Parse(time.DateTime, *dateTime)
	return err
}
