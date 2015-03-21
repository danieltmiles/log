package log

import (
	//"fmt"
	//"os"
	"sync"
	//"time"
)

/*
const (
	DEBUG              = 0
	INFO               = 1
	NOTICE             = 2
	WARNING            = 3
	ERROR              = 4
	FATAL              = 5
	defaultBufferLevel = WARNING
)
*/

var (
	tag            string
	mu             sync.Mutex
	hostname       string
	numNetWorkers  = 5
	numFileWorkers = 1
	numConnections = 1
)

/*
func init() {
	tag = os.Args[0]
	hostname, _ := os.Hostname()
}

func SetTag(t string) {
	tag = t
}
*/
