package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
	"time"
)

var (
	TIMELOG_FILE = path.Join(os.Getenv("HOME"), "timelog.txt")
)

func main() {
	var report = flag.String("report", "none", "Type of report you want to generate, options are full, projects ..")
	flag.Parse()
	//if report is none it means we are logging a task
	if *report == "none" {
		logTask()
	} else {
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
