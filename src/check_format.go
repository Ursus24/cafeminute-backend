package main

import (
	"strconv"
	"strings"
)

func checkDate(d string) bool {
	correct := true

	s := strings.Split(d, ("."))

	day, _ := strconv.ParseInt(s[0], 10, 0)
	month, _ := strconv.ParseInt(s[1], 10, 0)
	year, _ := strconv.ParseInt(s[2], 10, 0)

	if day > 31 {
		correct = false
	}

	if month > 12 {
		correct = false
	}

	if year < 2023 {
		correct = false
	}

	return correct
}

func checkTime(d string) bool {
	correct := true

	s := strings.Split(d, (":"))

	hour, _ := strconv.ParseInt(s[1], 10, 0)
	minutes, _ := strconv.ParseInt(s[2], 10, 0)

	if hour > 24 {
		correct = false
	}

	if minutes > 60 {
		correct = false
	}

	return correct
}
