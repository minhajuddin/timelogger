package main

// Contains code related to saving logs and reading logs

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

var (
	TIMELOG_FILE = path.Join(os.Getenv("HOME"), ".timelog.txt")
)

const (
	AVG_LINE_LENGTH  int64 = 50
	MAX_LOGS_PER_DAY int64 = 20
)

func logTask(msg string) {
	logFile, err := os.OpenFile(TIMELOG_FILE, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal("Unable to open timelog file:", err)
	}
	defer logFile.Close()

	//this is to stay consistent with the gtimelog format
	msg = strings.Replace(msg, "::", "**", 1)
	op := time.Now().Format("2006-01-02 15:04: ") + msg + "\n"
	_, err = logFile.WriteString(op)
	if err != nil {
		log.Fatal("Failed to write:", err)
	}
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

func filterLogs(logs []Log, predicate func(*Log) bool) []Log {
	oplogs := make([]Log, 0, len(logs))
	for _, l := range logs {
		if predicate(&l) {
			oplogs = append(oplogs, l)
		}
	}
	return oplogs
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

func getLatestLogs(n int64) {

	//TODO: this should probably be moved to main.go
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

//TODO Need the ability to seek further back if we have
//lesser lines than the number needed
//This method may return more logs than you need
//filter appropriately
//this code tries to get the last n lines from the file
//it is accurate most of the times, sometimes it might not
//be able to get n lines if the line size is large, in these
func readLatestLogs(n int64) []Log {
	fd, err := os.Open(TIMELOG_FILE)
	if err != nil {
		log.Fatal("Failed to open the timelog file for reading: ", TIMELOG_FILE, err)
	}
	defer fd.Close()

	//seek to these many lines from the back of the file
	_, _ = fd.Seek(-(AVG_LINE_LENGTH * n), 2)
	bytes, _ := ioutil.ReadAll(fd)
	lines := strings.Split(string(bytes), "\n")
	//we want to skip the first line as it might be read from the middle
	lines = lines[1:]
	return parseLines(lines)
}

//utility functions
func roundOffToDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
