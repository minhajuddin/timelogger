package main

import (
	"time"
)

type Filterer interface {
	Filter(LogReaderWriter) []Log
}

type LineFilter struct {
	Number int64
}

func (self *LineFilter) Filter(reader LogReaderWriter) []Log {
	return reader.Read(self.Number)
}

func getFilter(n int64) Filterer {
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

func filterLogs(logs []Log, predicate func(*Log) bool) []Log {
	oplogs := make([]Log, 0, len(logs))
	for _, l := range logs {
		if predicate(&l) {
			oplogs = append(oplogs, l)
		}
	}
	return oplogs
}

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
func roundOffToDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
