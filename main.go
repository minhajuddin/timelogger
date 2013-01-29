package main

//Contains code related to args parsing and dispatching

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var report = flag.String("report", "none", "Type of report you want to generate, options are full, projects ..")
	var days = flag.Int("d", -1, "Prints the logs for the last n days")
	var lineCount = flag.Int64("n", -1, "Prints the last n logs, prints 10 lines by default")
	flag.Parse()
	//if report is none it means we are logging a task
	switch {
	case len(os.Args) == 1 || *lineCount != -1:
		printLatestLogs(*lineCount)
	case *days != -1:
		printLogsForDays(*days)
	case *report == "none":
		logTask(strings.Join(os.Args[1:], " "))
	default:
		generateReport(*report)
	}
}

//TODO: To be moved to a better place
func generateReport(reportType string) {
	fmt.Println("Generating report:", reportType)
}

//utility functions
func roundOffToDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
