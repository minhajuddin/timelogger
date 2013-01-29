package main

// Contains code related to saving logs and reading logs

import (
	"io/ioutil"
	"os"
	"strings"
)

type LogReaderWriter interface {
	Read(n int) []Log
	Write(log *Log)
}

type TextReaderWriter struct {
	FilePath string
}

const (
	AVG_LINE_LENGTH  int64 = 50
	MAX_LOGS_PER_DAY int64 = 20
)

func (self *TextReaderWriter) Write(log *Log) {
	logFile, err := os.OpenFile(self.FilePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	handleError(err, "Unable to open timelog file")
	defer logFile.Close()

	op := log.End.Format("2006-01-02 15:04: ") + log.Text + "\n"

	_, err = logFile.WriteString(op)
	handleError(err, "Failed to write")
}

func (self *TextReaderWriter) Read(n int64) []Log {
	fd, err := os.Open(self.FilePath)
	handleError(err, "Failed to open the timelog file for reading:"+self.FilePath)
	defer fd.Close()
	//TODO Need the ability to seek further back if we have
	//lesser lines than the number needed
	// seek to these many lines from the back of the file
	_, _ = fd.Seek(-(AVG_LINE_LENGTH * n), 2)
	bytes, _ := ioutil.ReadAll(fd)
	lines := strings.Split(string(bytes), "\n")
	//we want to skip the first line as it might be read from the middle
	return parseLines(lines[1:])
}
