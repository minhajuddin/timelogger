package main

import (
	"fmt"
)

func printSummary(logs []Log) {
	for _, log := range logs {
		fmt.Println(log.String())
	}
}
