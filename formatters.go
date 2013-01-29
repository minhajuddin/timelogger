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
		fmt.Println(log.String())
	}
}

func getFormatter(formatter string) Formatter {
	return &PlainFormatter{}
}
