package log

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"
)

type Logger interface {
	Info(s ...string)
	Error(e error)
	Debug(s string)
}

type Level int

const (
	Error Level = 0
	Info  Level = 1
	Debug Level = 2
)

var levelTranslate = map[Level]string{
	Error: "ERROR",
	Info:  "INFO",
	Debug: "DEBUG",
}

type logger struct {
	level  Level
	logger *log.Logger
}

func NewLogger(l Level) Logger {
	log.SetFlags(0)
	return &logger{
		level:  l,
		logger: log.Default(),
	}
}

func (l *logger) log(level Level, s ...string) {
	l.logger.SetPrefix(fmt.Sprintf("%s [%s] ", time.Now().Format(time.RFC3339), levelTranslate[level]))

	l.logger.Print(strings.Join(s, " "))
}

func (l *logger) Info(s ...string) {
	if l.level < Info {
		return
	}
	l.log(Info, s...)
}

func (l *logger) Debug(s string) {
	if l.level < Debug {
		return
	}
	l.log(Debug, s)
}

func (l *logger) Error(e error) {
	var buf [16384]byte
	var usableStack string
	stack := buf[0:runtime.Stack(buf[:], false)]
	stackLines := strings.Split(string(stack), "\n")

	// remove linhas referentes às chamadas desta lib
	usableStack = strings.Join(stackLines[5:], "\n")
	l.log(Error, usableStack)
}
