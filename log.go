package log

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type Log struct {
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
	hostname     string
	tag          string
	levelStrings = []string{
		"FATAL",
		"ERROR",
		"WARNING",
		"NOTICE",
		"INFO",
		"DEBUG",
	}
)

func init() {
	hostname, _ = os.Hostname()
	tag = os.Args[0]
}

func New(writer io.Writer, threshold Level) *Log {
	return &Log{
		tag:       tag,
		threshold: threshold,
		writer:    writer,
	}
}

func (log *Log) SetTag(t string) {
	log.tag = t
}

func (log *Log) Debug(args ...interface{}) {
	log.write(Debug, args)
}

func (log *Log) Info(msg string) {
	log.write(Info, msg)
}

func (log *Log) Info(args ...interface{}) {
	log.write(Info, args)
}

func (log *Log) Warning(msg string) {
	log.write(Warning, msg)
}

func (log *Log) Notice(args ...interface{}) {
	log.write(Notice, args)
}

func (log *Log) Warning(args ...interface{}) {
	log.write(Warning, args)
}

func (log *Log) Error(args ...interface{}) {
	log.write(Error, args)
}

func (log *Log) Fatal(args ...interface{}) {
	log.write(Fatal, args)
	os.Exit(1)
}

func (log *Log) write(level Level, args ...interface{}) {
	if level > log.threshold {
		return
	}

	timestamp := time.Now().Format(time.RFC3339)

	log.mu.Lock()
	defer log.mu.Unlock()

	fmt.Fprintf(log.writer, "%s %s %s[%d]: %s %s\n",
		timestamp, hostname, log.tag, os.Getpid(), level, args)
}

func (level Level) String() string {
	return levelStrings[level]
}
