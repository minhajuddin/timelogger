package main

import (
	"fmt"
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
	return fmt.Sprintf("\t%s\t%5.2f - %s", self.Start.Format("02/01"), self.Duration().Hours(), self.Text)
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
	//convert it to local time
	t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, time.Local)
	return &Log{
		Text:    projectToken,
		End:     t,
		Project: project,
		Task:    task,
		Subtask: subtask,
	}
}

func parseLines(lines []string) []Log {
	logs := make([]Log, 0, 10)
	if len(lines) < 1 {
		return logs
	}
	//first pass, parse the end times
	for i := 0; i < len(lines); i++ {
		//skip empty lines
		//TODO should we treat empty lines like gtimelog?
		if len(lines[i]) > 0 {
			logs = append(logs, *parseLine(lines[i]))
		}
	}
	//second pass, compute the durations
	for i := 1; i < len(logs); i++ {
		logs[i].Start = logs[i-1].End
	}
	//cannot compute the summary for the first log entry
	return logs[1:]
}

func NewLog(message string) *Log {
	//this is to stay consistent with the gtimelog format
	return &Log{
		Text: strings.Replace(message, "::", "**", 1),
		End:  time.Now(),
	}
}
