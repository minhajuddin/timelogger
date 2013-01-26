package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

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
		//skip empty lines
		//TODO should we treat empty lines like gtimelog?
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

//TODO Need the ability to seek further back if we have
//lesser lines than the number needed
//This method may return more logs than you need
//filter appropriately
//this code tries to get the last n lines from the file
//it is accurate most of the times, sometimes it might not
//be able to get n lines if the line size is large, in these
func readLatestLogs(n int64) []*Log {
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
