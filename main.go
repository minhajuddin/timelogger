package main

import (
	"flag"
	"os"
)

func main() {
	formatterArg := flag.String("f", "plain", "Formatter for the output")
	//var days = flag.Int("d", -1, "Prints the logs for the last n days")
	//var lineCount = flag.Int64("n", -1, "Prints the last n logs, prints 10 lines by default")
	//flag.Parse()

	//create a io.Writer
	writer := os.Stdout
	//create a formatter with this writer
	formatter := getFormatter(*formatterArg)
	//run the query to get logs
	logs := getLogs()
	//pass the query through the formatter
	formatter.Format(logs, writer)
}
