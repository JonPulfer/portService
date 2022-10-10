package adapter

import (
	"github.com/JonPulfer/portsService/pb/message/port"
)

// Another option for receiving the port definition is via a message or pubsub service
// to stream the data.

type MessageQueue interface {
	Consume() (Message, error)
}

type Message struct {
	ID      string
	Type    string
	Payload *port.SerialisedPort
}
