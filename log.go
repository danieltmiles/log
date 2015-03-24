package log

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type Logger struct {
	mu        sync.Mutex
	tag       string
	threshold Level
	writer    io.Writer
}

type Level int8

const (
	None Level = iota - 1
	Fatal
	Error
	Warning
	Notice
	Info
	Debug
)

var (
	levelStrings = []string{
		"FATAL",
		"ERROR",
		"WARNING",
		"NOTICE",
		"INFO",
		"DEBUG",
	}
)

func New(writer io.Writer, threshold Level) *Logger {
	return &Logger{
		writer:    writer,
		threshold: threshold,
	}
}

func (log *Logger) SetTag(t string) {
	log.tag = t
}

func (log *Logger) Debug(msg string) {
	log.write(Debug, msg)
}

func (log *Logger) Info(msg string) {
	log.write(Info, msg)
}

func (log *Logger) Notice(msg string) {
	log.write(Notice, msg)
}

func (log *Logger) Warning(msg string) {
	log.write(Warning, msg)
}

func (log *Logger) Error(msg string) {
	log.write(Error, msg)
}

func (log *Logger) Fatal(msg string) {
	log.write(Fatal, msg)
	os.Exit(1)
}

func (log *Logger) write(level Level, msg string) {
	if level > log.threshold {
		return
	}

	timestamp := time.Now().Format(time.RFC3339)
	hostname, _ := os.Hostname()

	log.mu.Lock()
	defer log.mu.Unlock()

	fmt.Fprintf(log.writer, "%s %s %s[%d]: %s %s\n",
		timestamp, hostname, log.tag, os.Getpid(), level, msg)
}

func (level Level) String() string {
	return levelStrings[level]
}
