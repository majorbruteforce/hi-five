package broadcast

import (
	"errors"
)

// Encoding rules:
//
// All codes must be three digit long.
//
// The first digit must represent the direction of flow. (1: cs, 2: sc)
//
// The second digit must represent the nature of event. (0: server based, 1: session based)
//
// The final digit must represent the type of specific event.

const (
	//EventGetOnlineCount requests count of online clients.
	// client to server, server
	EventGetOnlineCount int = 100
	// EventRequestMatch indicates a request to find a match.
	// client to server, session
	EventReqMatch int = 110
	// EventReqSessionEnd indicates a request to end current session gracefully.
	// client to server, session
	EventReqSessionEnd int = 111
	// EventSendMessage indicates an message to be directed to the reciever.
	// client to server, session
	EventMatchmakingInProgress int = 112
	// EventMatchmakingInProgress indicates a request to find a match.
	// server to client, session
	EventSessionEnded int = 210
	// EventNewSession indicates a successful session creation.
	// server to client, session
	EventNewSession int = 211
	// EventSessionEnded indicates that the current session was ended.
	// server to client, session
	EventSendMessage int = 212
	// EventNewMessage indicates an incoming message from the server to the client.
	// server to client, session
	EventNewMessage int = 213
)

var (
	ErrEventNotSupported = errors.New("this event type is not supported")
)

type Event struct {
	Type    int         `json:"type"`
	Payload interface{} `json:"payload"`
}
