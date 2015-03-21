package log

import (
	"fmt"
	"os"
	"testing"

	. "github.com/franela/goblin"
	"github.com/monsooncommerce/mockstomp"
	. "github.com/onsi/gomega"
)

func TestLogger(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Logger Object Tests", func() {
		var (
			filePath       = "logFile.log"
			bufferDepth    int
			numConnections int
			username       string
			password       string
			queueName      string
			port           int
		)

		InitializeQueueConnection = func(host, username, password string, port int) (stompConnectioner, error) {
			return &mockstomp.MockStompConnection{}, nil
		}

		g.BeforeEach(func() {
			err := os.Remove("logFile.log")
			if err != nil {
				fmt.Printf("couldn't delete file: %v\n", err.Error())
			}
		})

		g.It("should initialize Logger object correctly", func() {
			username = "user"
			password = "pass"
			queueName = "myQueue"
			hostname = "queuehost.com"
			port = 1000
			filePath = "logFile.log"
			bufferDepth = 38000000
			numConnections = 50

			testLogger := NewLogger(hostname, username, password, filePath, queueName, port, bufferDepth, numConnections)

			Expect(cap(testLogger.fileChan)).To(Equal(bufferDepth))
			Expect(cap(testLogger.netChan)).To(Equal(bufferDepth))
			Expect(cap(testLogger.connectionChan)).To(Equal(numConnections))
		})
	})
}
