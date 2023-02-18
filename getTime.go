package main

import (
	"strconv"
	"time"
)

func getStartTime() string {
	loc, _ := time.LoadLocation("Local")
	startTime := time.Now()
	startTime = time.Date(startTime.Year(), startTime.Month(), startTime.Day(), startTime.Hour(), 00, 00, 0, loc)
	startTime_in_unix := startTime.UnixMilli()
	startTime_string := strconv.Itoa(int(startTime_in_unix))
	return startTime_string
}

func getEndTime() string {
	loc, _ := time.LoadLocation("Local")
	endTime := time.Now()
	endTime = time.Date(endTime.Year(), endTime.Month(), endTime.Day(), endTime.Hour(), 00, 01, 0, loc)
	endTime_in_unix := endTime.UnixMilli()
	endTime_string := strconv.Itoa(int(endTime_in_unix))
	return endTime_string
}
