package main

import (
	"fmt"
	"io"
	"time"
)

//Project | Time spent in hours | Date
//TL      | 8.33                | 2012-10-23
//TL      | 6.33                | 2012-10-24
//Learn   | 8.3                 | 2012-10-24
//Break   | 3                   | 2012-10-25

type ProjectSummaryFormatter struct{}

type ProjectSummaryLog struct {
	Project  string
	Tasks    []string
	Duration time.Duration
	Date     time.Time
}

func projectSummaryLogHash(l *Log) string {
	return l.End.Format("20060102:") + l.Project
}

func (self ProjectSummaryFormatter) Format(logs []Log, writer io.Writer) {
	summary := make(map[string]*ProjectSummaryLog)

	//map of date:project
	for _, log := range logs {
		hash := projectSummaryLogHash(&log)
		if _, ok := summary[hash]; !ok {
			summary[hash] = &ProjectSummaryLog{
				Project: log.Project,
				Date:    log.End,
			}
		}
		p := summary[hash]
		//TODO add current task to the task list
		p.Duration += log.Duration()
	}

	//TODO: add sorting
	for _, p := range summary {
		fmt.Printf("\t%s\t%5.2f - %s\n", p.Date.Format("2006-01-02"), p.Duration.Hours(), p.Project)
	}

}
