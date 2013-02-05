package main

import (
	"fmt"
	"io"
	"sort"
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

type ProjectSummaryLogs []ProjectSummaryLog

func (s ProjectSummaryLogs) Len() int           { return len(s) }
func (s ProjectSummaryLogs) Less(i, j int) bool { return s[j].Date.After(s[i].Date) }
func (s ProjectSummaryLogs) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

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

	pslogs := make(ProjectSummaryLogs, 0, len(summary))

	for _, p := range summary {
		pslogs = append(pslogs, *p)
	}

	sort.Sort(pslogs)

	//TODO: add sorting
	for _, p := range pslogs {
		fmt.Printf("\t%s\t%5.2f - %s\n", p.Date.Format("2006-01-02"), p.Duration.Hours(), p.Project)
	}

}
