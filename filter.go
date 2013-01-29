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

//TODO: change all int64 to int
func getFilter(n int64, days int) Filterer {
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

//========================================
//========================================

//TODO:
//	- n number
//  -since date
//  -grep date

func readLatestLogs(n int64) []Log {
	return make([]Log, 10)
}

////if report is none it means we are logging a task
//switch {
//case len(os.Args) == 1 || *lineCount != -1:
//printLatestLogs(*lineCount)
//case *days != -1:
//printLogsForDays(*days)
//case *report == "none":
//logTask(strings.Join(os.Args[1:], " "))
//default:
//generateReport(*report)
//}
func getLogs() []Log {
	return readLatestLogs(10)
}

func printSummary(logs []Log) {
}

func printLogsForDays(days int) {
	tillDate := time.Now().AddDate(0, 0, -1*(days-1))
	tillDate = roundOffToDate(tillDate)
	logs := readLatestLogs(int64(days) * MAX_LOGS_PER_DAY)
	logs = filterLogs(logs, func(l *Log) bool { return l.End.After(tillDate) })
	printSummary(logs)
}

func printLatestLogs(n int64) {

	//this is when no option is given 
	if n == -1 {
		n = 10
	}
	logs := readLatestLogs(n)
	lindex := int64(len(logs)) - n
	//if we have less lines than the lineCount, show all the lines
	if lindex < 0 {
		lindex = 0
	}
	printSummary(logs[lindex:])
}

//utility functions
