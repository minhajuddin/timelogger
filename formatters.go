package main

import (
	"fmt"
	"io"
	"time"
)

type Formatter interface {
	Format(logs []Log, writer io.Writer)
}

type PlainFormatter struct{}

func (self PlainFormatter) Format(logs []Log, writer io.Writer) {
	for _, log := range logs {
		fmt.Printf("\t%2d:%02d - %s\n", int(log.Duration().Hours()), int(log.Duration().Minutes())%60, log.Text)
	}
}

func getFormatter(formatter string) Formatter {
	if formatter == "project-summary" {
		return &ProjectSummaryFormatter{}
	}
	return &PlainFormatter{}
}

//Project | Time spent in hours | Date
//TL      | 8.33                | 2012-10-23
//TL      | 6.33                | 2012-10-24
//Learn   | 8.3                 | 2012-10-24
//Break   | 3                   | 2012-10-25

type ProjectSummaryFormatter struct{}

type ProjectSummaryLog struct {
	Project string
	Tasks   []string
	Hours   time.Duration
	Date    time.Time
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
		p.Hours += log.Duration()
	}

	for _, p := range summary {
		fmt.Printf("%4s%4s\t%4v\n", p.Project, p.Date.Format("2006-01-02"), p.Hours)
	}

}
