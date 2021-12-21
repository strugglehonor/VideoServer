package log

import (
	"sync"
)

type logger struct {
	opt  *options
	entryPool *sync.Pool
	mu  sync.Mutex
}

var std = NewLogger()

func NewLogger(opts ...Option) *logger {
	return &logger{
		opt: initOptions(opts...),
		entryPool: &sync.Pool{New: func() interface{} {
			return new(Entry)
		}},
	}
}