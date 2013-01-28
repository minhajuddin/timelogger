package main

import (
	"testing"
	"time"
)

func assertEqual(t *testing.T, expected, actual interface{}, message string) {
	if expected != actual {
		t.Log(message)
		t.Errorf("Equality assertion failed: %v != %v", expected, actual)
	}
}

func TestParse(t *testing.T) {
	log := parseLine("2013-01-27 18:13: timelogger dev")

	if log.Text != "timelogger dev" {
		t.Error("Unable to parse log text")
	}

	if log.Project != "timelogger" {
		t.Error("Unable to parse project")
	}

	if log.Task != "dev" {
		t.Error("Unable to parse task")
	}

	if log.Subtask != "" {
		t.Error("Unable to parse subtask")
	}

	assertEqual(t, time.Date(2013, 01, 27, 18, 13, 0, 0, time.UTC), log.End, "")

}
