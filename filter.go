package main

import (
	"time"
)

type Filterer interface {
	Filter(LogReaderWriter) []Log
}

func filterLogs(logs []Log, predicate func(*Log) bool) []Log {
	oplogs := make([]Log, 0, len(logs))
	for _, l := range logs {
		if predicate(&l) {
			oplogs = append(oplogs, l)
		}
	}
	return oplogs
}

//TODO:
//  -grep date

//TODO: change all int64 to int
func getFilter(n int64, days int, since string) Filterer {
	if since != "" {
		date, err := time.Parse("2006-01-02", since)
		handleError(err, "Incorrect date format. Should be input as 2012-02-22")
		return &DateFilter{Date: date}
	}
	if days != 0 {
		date := time.Now().AddDate(0, 0, -1*(int(days)-1))
		date = roundOffToDate(date)
		return &DateFilter{Date: date}
	}
	if n == 0 {
		n = 10
	}
	return &LineFilter{Number: n}
}
