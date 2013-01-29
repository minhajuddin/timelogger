package main

//Contains code related to args parsing and dispatching

import (
	"flag"
	"fmt"
	"os"
	_ "strings"
	"time"
)

func main() {
	formatterArg := flag.String("f", "plain", "Formatter for the output")
	//var days = flag.Int("d", -1, "Prints the logs for the last n days")
	//var lineCount = flag.Int64("n", -1, "Prints the last n logs, prints 10 lines by default")
	//flag.Parse()

	//create a io.Writer
	//create a formatter with this writer
	formatter := getFormatter(*formatterArg)
	//run the query to get logs
	//TODO: remove hardcoded number
	logs := readLatestLogs(10)
	//pass the query through the formatter
	formatter.Format(logs, os.Stdout)

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
}

//TODO: To be moved to a better place
func generateReport(reportType string) {
	fmt.Println("Generating report:", reportType)
}

//utility functions
func roundOffToDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
