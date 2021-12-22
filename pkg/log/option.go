package log

import (
	"io"
	"os"
)

type level uint8

const (
	DebugLevel level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

var LevelNameMapping = map[level]string {
	DebugLevel: "Debug",
	InfoLevel: "Info",
	WarnLevel: "Warn",
	ErrorLevel: "Error",
	PanicLevel: "Panic",
	FatalLevel: "Fatal",
}

type Option func(*options)

type options struct {
	Level  level
	StdLevel level
	Format Formatter
	IsColor bool
	output  io.Writer
}

// stdLevel level, format Format, isColor bool
func initOptions(opts ...Option) (o *options) {
	o = &options{}
	for _, opt := range opts {
		opt(o)
	}
	
	// if o.Format == nil {
	// 	o.Format = &Formatter{}
	// }

	if o.output == nil {
		o.output = os.Stderr
	}

	return o
}

func withOutput(output io.Writer) Option {
	return func(o *options) {
		o.output = output
	}
}

func withColor() Option {
	return func(o *options) {
		o.IsColor = true
	}
}

// func withFormat(formation Format) Option {
// 	return func(o *options) {
// 		o.Format = &formation
// 	}
// }

func withStdLevel(stdLevel level) Option {
	return func(o *options) {
		o.StdLevel = stdLevel
	}
}
