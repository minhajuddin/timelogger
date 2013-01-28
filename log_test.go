package main

import (
	"github.com/remogatto/prettytest"
	"testing"
	"time"
)

//start setup
type testSuite struct {
	prettytest.Suite
}

func TestRunner(t *testing.T) {
	prettytest.RunWithFormatter(
		t,
		new(prettytest.TDDFormatter),
		new(testSuite),
	)
}

var l = parseLine("2013-01-27 18:13: timelogger dev")

func (t *testSuite) TestTextParse() {
	t.Equal(l.Text, "timelogger dev")
}
func (t *testSuite) TestProjectParse() {
	t.Equal(l.Project, "timelogger")
}
func (t *testSuite) TestTaskParse() {
	t.Equal(l.Task, "dev")
}
func (t *testSuite) TestSubtaskParse() {
	t.Equal(l.Subtask, "")
}

func (t *testSuite) TestDateParse() {
	t.Equal(l.End, time.Date(2013, 01, 27, 18, 13, 0, 0, time.UTC))
}

func (t *testSuite) TestParseLines() {
	logs := parseLines([]string{
		"2013-01-26 14:34: timelogger warmup",
		"2013-01-26 14:58: break",
	})

	t.Equal(len(logs), 1)
	t.Equal(logs[0].Start, time.Date(2013, 01, 26, 14, 34, 0, 0, time.UTC))
}
