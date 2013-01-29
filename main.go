package main

import (
	"flag"
	"os"
	"path"
	"strings"
)

var message = flag.String("m", "", "Message to be logged")
var formatterArg = flag.String("f", "plain", "Formatter for the output")
var days = flag.Int("d", -1, "Prints the logs for the last n days")

//var lineCount = flag.Int64("n", -1, "Prints the last n logs, prints 10 lines by default")

func main() {
	flag.Parse()
	logReaderWriter := &TextReaderWriter{FilePath: path.Join(os.Getenv("HOME"), ".timelog.txt")}

	//if no flags were passed write a log
	if noFlags() && len(os.Args) > 1 {
		log := NewLog(strings.Join(os.Args[1:], " "))
		logReaderWriter.Write(log)
	}

	////create a io.Writer
	//writer := os.Stdout
	////create a formatter with this writer
	//formatter := getFormatter(*formatterArg)
	////filter the logs using the current filter
	//filter := getFilter()
	//logs := filter.Filter()
	////pass the query through the formatter
	//formatter.Format(logs, writer)
}

func noFlags() bool {
	return len(flag.Args()) == len(os.Args)-1
}
