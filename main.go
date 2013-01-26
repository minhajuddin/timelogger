package main

import (
	"flag"
	"fmt"
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

func main() {
	var report = flag.String("report", "none", "Type of report you want to generate, options are full, projects ..")
	var days = flag.Int64("d", -1, "Prints the logs for the last n days")
	var lineCount = flag.Int64("n", -1, "Prints the last n logs, prints 10 lines by default")
	flag.Parse()
	//if report is none it means we are logging a task
	switch {
	case len(os.Args) == 1 || *lineCount != -1:
		printLatestLogs(*lineCount)
	case *days != -1:
		printLogsForDays(*days)
	case *report == "none":
		logTask()
	default:
		generateReport(*report)
	}
}

func logTask() {
	logFile, err := os.OpenFile(TIMELOG_FILE, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal("Unable to open timelog file:", err)
	}
	defer logFile.Close()

	msg := strings.Join(os.Args[1:], " ")
	//this is to stay consistent with the gtimelog format
	msg = strings.Replace(msg, "::", "**", 1)
	op := time.Now().Format("2006-01-02 15:04: ") + msg + "\n"
	_, err = logFile.WriteString(op)
	if err != nil {
		log.Fatal("Failed to write:", err)
	}
}

func generateReport(reportType string) {
	fmt.Println("Generating report:", reportType)
}

func printLogsForDays(days int64) {
	printLatestLogs(days * MAX_LOGS_PER_DAY)
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

func printSummary(logs []*Log) {
	for _, log := range logs {
		fmt.Println(log.String())
	}
}
