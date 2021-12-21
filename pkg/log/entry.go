package log

import (
	"bytes"
	"runtime"
	"time"
)

// 日志写入格式，日志配置和日志内容，日志的堆栈，所在的行，文件，
type Entry struct {
	logger *logger
	file  string
	line  string
	Func  string
	Format  string
	level  level
	time  time.Time
	buffer bytes.Buffer
	args  []interface{}
}

func NewEntry(l *logger) *Entry {
	return &Entry{
		logger: l,
	}
}

func (e Entry) Write(format string, level level, args ...interface{}) {
	if e.logger.opt.Level > level {
		return 
	}
	e.Format = format
	e.args = args
	e.level = level
	runtime.Caller(2)
}