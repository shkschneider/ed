package main

import(
	"bytes"
	"fmt"
	"io"
	// "os"
	"time"

	pretty "github.com/k0kubun/pp/v3"
)

type Log struct {
	writer	io.Writer
	level	LogLevel
	colors	bool
}

var	log *Log
var pp *pretty.PrettyPrinter

type LogLevel struct {
	Tag string
	Color uint8
}

var (
	LogLevelDebug   = LogLevel { "DBG", 34 }
	LogLevelInfo    = LogLevel { "INF", 32 }
	LogLevelWarning = LogLevel { "WRN", 33 }
	LogLevelError   = LogLevel { "ERR", 31 }
	LogLevelFatal   = LogLevel { "WTF", 35 }
)

func NewLog(writer io.Writer, level LogLevel, colors bool) {
	log = &Log { writer, level, colors } // io.MultiWriter(writer, os.Stderr)
	pp = pretty.New()
	pp.SetColoringEnabled(false)
}

func (l Log) log(lvl LogLevel, any string) {
	var buffer bytes.Buffer
	if l.colors {
		buffer.WriteString(fmt.Sprintf("\033[%dm", lvl.Color))
	}
	buffer.WriteString(fmt.Sprintf("%s [%3s]",
		time.Now().Format(fmt.Sprintf("%s %s.000", time.DateOnly, time.TimeOnly)), lvl.Tag))
	if l.colors {
		buffer.WriteString(fmt.Sprintf("\033[0m"))
	}
	buffer.WriteString(fmt.Sprintf(" %s\n", any))
	l.writer.Write(buffer.Bytes())
}

func (l Log) Debug(many ...interface{}) {
	for _, any := range many {
		l.log(LogLevelDebug, pp.Sprint(any))
	}
}

func (l Log) Info(many ...interface{}) {
	for _, any := range many {
		l.log(LogLevelInfo, pp.Sprint(any))
	}
}

func (l Log) Warning(many ...interface{}) {
	for _, any := range many {
		l.log(LogLevelWarning, pp.Sprint(any))
	}
}

func (l Log) Error(many ...interface{}) {
	for _, any := range many {
		l.log(LogLevelError, pp.Sprint(any))
	}
}

func (l Log) Fatal(many ...interface{}) {
	for _, any := range many {
		l.log(LogLevelFatal, pp.Sprint(any))
	}
	if len(many) == 0 {
		panic("!!!")
	} else {
		panic(any(many))
	}
}
