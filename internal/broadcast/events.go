package broadcast

import (
	"errors"
)

const (
	// EventRequestMatch indicates a request to find a match
	EventRequestMatch int = 100
	// EventMatchmakingInProgress indicates a request to find a match
	EventMatchmakingInProgress int = 101
	// EventNewSession indicates a successful session creation
	EventNewSession int = 102
	//EventEndSession indicates a request to end current session gracefully
	EventEndSession int = 103
	// EventSendMessage indicates an message to be directed to the reciever
	EventSendMessage int = 200
	// EventNewMessage indicates an incoming message from the server to the client
	EventNewMessage int = 201
	//EventGetOnlineCount requests count of online clients
	EventGetOnlineCount int = 400
)

var (
	ErrEventNotSupported = errors.New("this event type is not supported")
)

type Event struct {
	Type    int    `json:"type"`
	Payload string `json:"payload"`
}
