package main

import (
	"flag"
	"os"
	"path"
	"strings"
)

var n = flag.Int64("n", 0, "Prints the last n logs, prints 10 lines by default")
var days = flag.Int("d", 0, "Prints the logs for the last n days")
var formatterArg = flag.String("f", "plain", "Formatter for the output")

func main() {
	flag.Parse()
	logReaderWriter := &TextReaderWriter{FilePath: path.Join(os.Getenv("HOME"), ".timelog.txt")}

	//if no flags were passed write a log
	if noFlags() && len(os.Args) > 1 {
		log := NewLog(strings.Join(os.Args[1:], " "))
		logReaderWriter.Write(log)
		return
	}

	//create a io.Writer
	writer := os.Stdout
	//create a formatter with this writer
	formatter := getFormatter(*formatterArg)
	//filter the logs using the current filter
	filter := getFilter(*n, *days)
	logs := filter.Filter(logReaderWriter)
	////pass the query through the formatter
	formatter.Format(logs, writer)
}

func noFlags() bool {
	return len(flag.Args()) == len(os.Args)-1
}
