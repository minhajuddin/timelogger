package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

var (
	TIMELOG_FILE = path.Join(os.Getenv("HOME"), "timelog.txt")
)

const (
	AVG_LINE_LENGTH int64 = 50
)

func main() {
	var report = flag.String("report", "none", "Type of report you want to generate, options are full, projects ..")
	var lineCount = flag.Int64("lines", 10, "Prints the last n logs")
	flag.Parse()
	//if report is none it means we are logging a task
	switch {
	case len(os.Args) == 1:
		printLatestLogs(*lineCount)
	case *report == "none":
		logTask()
	default:
		generateReport(*report)
	}
}

func logTask() {
	logFile, err := os.OpenFile(TIMELOG_FILE, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("Unable to open timelog file:", err)
		return
	}
	defer logFile.Close()

	op := time.Now().Format("2006-01-02 15:04: ") + strings.Join(os.Args[1:], " ") + "\n"
	_, err = logFile.WriteString(op)
	if err != nil {
		fmt.Println("Failed to write:", err)
	}
}

func generateReport(reportType string) {
	fmt.Println("Generating report:", reportType)
}

//this code tries to get the last n lines from the file
//it is accurate most of the times, sometimes it might not
//be able to get n lines if the line size is large, in these
func printLatestLogs(lineCount int64) {
	fd, err := os.Open("/home/minhajuddin/.gtimelog/timelog.txt")
	if err != nil {
		log.Fatal("Failed to open the timelog file for reading: ", TIMELOG_FILE, err)
	}

	_, _ = fd.Seek(-(AVG_LINE_LENGTH * lineCount), 2)
	bytes, _ := ioutil.ReadAll(fd)
	lines := strings.Split(string(bytes), "\n")
	//we want to skip the first line as it might be read from the middle
	lines = lines[1:]
	lindex := int64(len(lines)) - lineCount - 2
	printSummary(lines[lindex:])
}

func printSummary(logs []string) {
	//TODO: make this a filter which shows the actual time spent on the tasks
	fmt.Println(strings.Join(logs, "\n"))
}
