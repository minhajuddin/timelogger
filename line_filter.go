package main

type LineFilter struct {
	Number int64
}

func (self *LineFilter) Filter(reader LogReaderWriter) []Log {
	logs := reader.Read(self.Number)
	lindex := int64(len(logs)) - self.Number
	return logs[lindex:]
}
