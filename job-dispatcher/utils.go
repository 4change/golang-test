package main

import (
	"log"
	"strconv"
	"strings"
	"time"
)

type CustomerTime struct {
	Time  time.Time
	Valid bool
}

func (t *CustomerTime) UnmarshalJSON(data []byte) error {
	timeStr, err := strconv.Unquote(string(data))
	if err != nil {
		log.Println(err)
	}
	timeStr = strings.Replace(timeStr, " ", "T", 1)

	newTime, err := time.Parse("2006-01-02T15:04:05", timeStr)
	if err != nil {
		log.Println(err)
		*t = CustomerTime{}
		return nil
	}
	*t = CustomerTime{Time: newTime, Valid: true}
	return nil
}

func (ct CustomerTime) Value() interface{} {
	if !ct.Valid {
		return nil
	}
	return ct.Time
}

func (ct CustomerTime) FormatToDay() string {
	return ct.Time.Format("2006-01-02")
}

func (ct CustomerTime) FormatToHour() string {
	return ct.Time.Format("2006-01-02 15:00:00")
}

func (ct CustomerTime) String() string {
	return ct.Time.Format("2006-01-02 15:04:05")
}
