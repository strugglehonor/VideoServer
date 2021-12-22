package log

import (
	"bytes"
	"runtime"
	"time"
)

var FmtEmptySeparate = ""

// 日志写入格式，日志配置和日志内容，日志的堆栈，所在的行，文件，
type Entry struct {
	logger *logger
	file  string
	line  int
	Func  string
	Format  string
	level  level
	time  time.Time
	buffer bytes.Buffer
	args  []interface{}
}

func (l *logger) entry() *Entry {
	return &Entry{
		logger: l,
	}
}

func (e *Entry) write(level level, format string, args ...interface{}) {
	if e.logger.opt.Level > level {
		return 
	}

	e.Format = format
	e.args = args
	e.level = level
	e.time = time.Now()

	if pc, file, line, ok := runtime.Caller(2); ok {
		e.file, e.line = file, line
		e.Func = runtime.FuncForPC(pc).Name()
	}else {
		e.file, e.Func, e.line = "???", "???", 0
	}
}

func (e *Entry) format() {
	e.logger.opt.Format.Format(e)
}

func (e *Entry) writer() {
	e.logger.mu.Lock()
	e.logger.opt.output.Write(e.buffer.Bytes())
	e.logger.mu.Unlock()
}

func (e *Entry) release() {
	e.Func, e.line, e.file, e.Format, e.args, e.level = "", 0, "", "", nil, 0
	e.buffer.Reset()
	e.logger.entryPool.Put(e)
}