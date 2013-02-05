package main

import (
	"fmt"
	"io"
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
