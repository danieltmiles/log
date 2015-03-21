package log

import (
	"log"
	"net"

	"github.com/gmallard/stompngo"
)

type stompConnectioner interface {
	Send(stompngo.Headers, string) error
}

var InitializeQueueConnection = func(host, username, password string, port int) (stompConnectioner, error) {
	stompNetConn, err := net.Dial("tcp", net.JoinHostPort(host, string(port)))
	if err != nil {
		log.Fatal("unable to dial broker: " + err.Error())
		return nil, err
	}
	headers := stompngo.Headers{
		"accept-version", "1.1",
		"login", username,
		"passcode", password,
		"host", host,
	}
	stompConnection, err := stompngo.Connect(stompNetConn, headers)
	if err != nil {
		log.Fatal("unable to connect to broker: " + err.Error())
		return nil, err
	}
	return stompConnection, nil
}

func sendToQueue(stompConnection stompConnectioner, queueName, msg string) {
	headers := stompngo.Headers{
		"persistent", "true",
		"destination", "/queue/" + queueName,
	}
	err := stompConnection.Send(headers, msg)
	if err != nil {
		log.Fatal("unable to send stomp message: " + err.Error())
	}
}
