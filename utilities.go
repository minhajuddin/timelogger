package main

import (
	"log"
	"time"
)

func handleError(err error, message string) {
	if err != nil {
		log.Fatal("ERR:", message, err)
	}
}

func roundOffToDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func roundFloat(n float64) int64 {
	if n > 0 {
		return int64(n + 0.5)
	}
	return int64(n - 0.5)
}

func out(args ...interface{}) {
	log.Println(args...)
}
