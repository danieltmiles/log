package log

import (
	"fmt"
	"os"
	"time"
)

type Logger struct {
	stompVersion string
	username     string
	password     string
	host         string
	queueName    string

	outFile *os.File

	fileChan       chan string
	netChan        chan string
	connectionChan chan stompConnectioner
}

func NewLogger(host, username, password, filePath, queueName string, port, bufferDepth, numConnections int) *Logger {
	l := &Logger{}

	if filePath == "stdout" {
		l.outFile = os.Stdout
	} else if filePath == "stderr" {
		l.outFile = os.Stderr
	} else {
		logfile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0444)
		if err != nil {
			fmt.Printf("couldn't open logfile: %v\n", err.Error())
		}
		l.outFile = logfile
	}

	l.host = host
	l.password = password
	l.queueName = queueName
	l.stompVersion = "1.1"
	l.username = username

	l.connectionChan = make(chan stompConnectioner, numConnections)
	l.fileChan = make(chan string, bufferDepth)
	l.netChan = make(chan string, bufferDepth)

	for i := 0; i < numNetWorkers; i++ {
		go l.netWorker()
	}
	for i := 0; i < numFileWorkers; i++ {
		go l.fileWorker()
	}
	for i := 0; i < numConnections; i++ {
		conn, err := InitializeQueueConnection(host, username, password, port)
		if err != nil {
			fmt.Printf("Couldn't connect to rabbit")
			break
		}
		l.connectionChan <- conn
	}
	return l
}

func (l *Logger) sendStompMessage(msg string) {
	select {
	case conn := <-l.connectionChan:
		sendToQueue(conn, l.queueName, msg)
	default:
	}
}

func (l *Logger) netWorker() {
	for {
		msg := <-l.netChan
		l.sendStompMessage(msg)
	}
}

func (l *Logger) fileWorker() {
	for {
		msg := <-l.fileChan
		l.outFile.Write([]byte(msg))
		l.outFile.Write([]byte("\n"))
	}
}

func (l *Logger) Debug(msg string) {
	l.publish("DEBUG", msg)
}

func (l *Logger) Error(msg string) {
	l.publish("ERROR", msg)
}

func (l *Logger) Fatal(msg string) {
	l.publish("FATAL", msg)
	os.Exit(1)
}

func (l *Logger) Info(msg string) {
	l.publish("INFO", msg)
}

func (l *Logger) Notice(msg string) {
	l.publish("NOTICE", msg)
}

func (l *Logger) Warning(msg string) {
	l.publish("WARNING", msg)
}

func (l *Logger) publish(logLevel, msg string) {
	//formattedMsg := format(logLevel, msg)
	timestamp := time.Now().Format(time.RFC3339)
	formattedMsg := fmt.Sprintf(
		"%s %s %s[%d]: %s %s\n",
		timestamp,
		hostname,
		tag,
		os.Getpid(),
		logLevel,
		msg,
	)
	select {
	case l.fileChan <- formattedMsg:
	default:
	}
	select {
	case l.netChan <- formattedMsg:
	default:
	}
}
