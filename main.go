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
	TIMELOG_FILE = path.Join(os.Getenv("HOME"), ".timelog.txt")
)

const (
	AVG_LINE_LENGTH int64 = 50
)

func main() {
	var report = flag.String("report", "none", "Type of report you want to generate, options are full, projects ..")
	var lineCount = flag.Int64("lines", -1, "Prints the last n logs, prints 10 lines by default")
	flag.Parse()
	//if report is none it means we are logging a task
	switch {
	case len(os.Args) == 1 || *lineCount != -1:
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
		log.Fatal("Unable to open timelog file:", err)
	}
	defer logFile.Close()

	op := time.Now().Format("2006-01-02 15:04: ") + strings.Join(os.Args[1:], " ") + "\n"
	_, err = logFile.WriteString(op)
	if err != nil {
		log.Fatal("Failed to write:", err)
	}
}

func generateReport(reportType string) {
	fmt.Println("Generating report:", reportType)
}

//this code tries to get the last n lines from the file
//it is accurate most of the times, sometimes it might not
//be able to get n lines if the line size is large, in these
func printLatestLogs(lineCount int64) {
	fd, err := os.Open(TIMELOG_FILE)
	if err != nil {
		log.Fatal("Failed to open the timelog file for reading: ", TIMELOG_FILE, err)
	}

	//this is when no option is when 
	if lineCount == -1 {
		lineCount = 10
	}

	_, _ = fd.Seek(-(AVG_LINE_LENGTH * lineCount), 2)
	bytes, _ := ioutil.ReadAll(fd)
	lines := strings.Split(string(bytes), "\n")
	//we want to skip the first line as it might be read from the middle
	lines = lines[1:]
	lindex := int64(len(lines)) - lineCount - 2
	//if we have less lines than the lineCount, show all the lines
	if lindex < 0 {
		lindex = 0
	}
	printSummary(lines[lindex:])
}

//TODO: change refs to pointers if it makes sense
type Log struct {
	Text    string
	Start   time.Time
	End     time.Time
	Project string
	Task    string
	Subtask string
}

func (self *Log) Duration() time.Duration {
	return self.End.Sub(self.Start)
}

func (self *Log) String() string {
	return fmt.Sprintf("\t%5.2f - %s", self.Duration().Hours(), self.Text)
}

//"2013-01-18 15:24: learn code-reading gostatic"
func parseLine(line string) *Log {
	tokens := strings.Split(line, ": ")
	timeToken := tokens[0]
	projectToken := tokens[1]
	projectTokens := strings.Fields(projectToken)

	var project, task, subtask string

	if len(projectTokens) > 2 {
		subtask = strings.Join(projectTokens[2:], " ")
	}
	if len(projectTokens) > 1 {
		task = projectTokens[1]
	}
	if len(projectTokens) > 0 {
		project = projectTokens[0]
	}
	t, _ := time.Parse("2006-01-02 15:04", timeToken)
	return &Log{
		Text:    projectToken,
		End:     t,
		Project: project,
		Task:    task,
		Subtask: subtask,
	}
}

func parseLines(lines []string) []*Log {
	logs := make([]*Log, 0, 10)
	//first pass, parse the end times
	for i := 0; i < len(lines); i++ {
		if len(lines[i]) > 0 {
			logs = append(logs, parseLine(lines[i]))
		}
	}
	//second pass, compute the durations
	for i := 1; i < len(logs); i++ {
		logs[i].Start = logs[i-1].End
	}
	//cannot compute the summary for the first log entry
	return logs[1:]
}

func printSummary(lines []string) {
	//TODO: make this a filter which shows the actual time spent on the tasks
	logs := parseLines(lines)
	for _, log := range logs {
		fmt.Println(log.String())
	}
}
