package main

type LineFilter struct {
	Number int64
}

func (self *LineFilter) Filter(reader LogReaderWriter) []Log {
	return reader.Read(self.Number)
}
