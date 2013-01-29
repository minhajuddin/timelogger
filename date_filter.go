package main

import (
	"time"
)

type DateFilter struct {
	Date time.Time
}

func (self *DateFilter) Days() int64 {
	today := roundOffToDate(time.Now())
	return roundFloat(today.Sub(self.Date).Hours() / 24)
}

func (self *DateFilter) Filter(reader LogReaderWriter) []Log {
	logs := reader.Read(self.Days() * MAX_LOGS_PER_DAY)
	logs = filterLogs(logs, func(l *Log) bool { return l.End.After(self.Date) })
	return logs
}
