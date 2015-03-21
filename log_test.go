package log

import (
	"fmt"
	"os"
	"testing"
	"time"

	. "github.com/franela/goblin"
	"github.com/monsooncommerce/mockstomp"
	. "github.com/onsi/gomega"
)

func TestLogger(t *testing.T) {
	g := Goblin(t)
	RegisterFailHandler(func(m string, _ ...int) { g.Fail(m) })

	g.Describe("Logger Object Tests", func() {
		var (
			filePath        = "logFile.log"
			bufferDepth     int
			numConnections  int
			username        string
			password        string
			queueName       string
			port            int
			stompConnection mockstomp.MockStompConnection
		)

		g.BeforeEach(func() {
			stompConnection = mockstomp.MockStompConnection{}
			err := os.Remove("logFile.log")
			if err != nil {
				fmt.Printf("couldn't delete file: %v\n", err.Error())
			}
			InitializeQueueConnection = func(host, username, password string, port int) (stompConnectioner, error) {
				stompConnection.Init()

				return &stompConnection, nil
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
			numConnections = 1

			testLogger := NewLogger(hostname, username, password, filePath, queueName, port, bufferDepth, numConnections)

			Expect(cap(testLogger.fileChan)).To(Equal(bufferDepth))
			Expect(cap(testLogger.netChan)).To(Equal(bufferDepth))
			Expect(cap(testLogger.connectionChan)).To(Equal(numConnections))

			Expect(testLogger.host).To(Equal(hostname))
			Expect(testLogger.password).To(Equal(password))
			Expect(testLogger.queueName).To(Equal(queueName))
		})

		g.It("should send messages to mock queue", func() {

			testLogger := NewLogger(
				"user", "pass", "myQueue", "queuehost.com",
				"logFile.log", 1000, 38000000, 1,
			)

			logLine := "Can't connect to server."
			testLogger.Debug(logLine)

			expectedMessage := mockstomp.MockStompMessage{
				Order:   0,
				Headers: []string{"persistent", "true", "destination", "/queue/logFile.log"},
				Message: fmt.Sprintf(
					"%v queuehost.com [%v]: DEBUG Can't connect to server.\n",
					time.Now().Format("2006-01-02T15:04:05-07:00"),
					os.Getpid(),
				),
			}

			time.Sleep(1 * time.Second)
			//stompConnection := <-testLogger.connectionChan

			Expect(len(stompConnection.MessagesSent)).To(Equal(1))
			msg := <-stompConnection.MessagesSent
			Expect(msg).To(Equal(expectedMessage))
		})
	})
}
