package broadcast

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/gorilla/websocket"

	"github.com/google/uuid"
)

type Client struct {
	ID          string
	DisplayName string
	connection  *websocket.Conn
	egress      chan Event
}

func NewClient(conn *websocket.Conn, displayName string) *Client {
	return &Client{
		ID:          uuid.New().String(),
		DisplayName: displayName,
		connection:  conn,
		egress:      make(chan Event),
	}
}

// readMessages constantly reads for incoming events from the connection.
// and routes them through routeEvent.
//
// It will only route the messages that are of event type client to server (< 200)
func (c *Client) readMessages(cm *ConnectionManager) {
	defer func() {
		cm.removeClient(c)
		// If client has existing session, remove it
		if _, ok := cm.sessions[c]; ok {
			cm.removeSession(c)
		}
	}()

	for {
		// ReadMessage is used to read the next message in queue
		// in the connection
		_, data, err := c.connection.ReadMessage()

		if err != nil {
			// If Connection is closed, we will Recieve an error here
			// We only want to log Strange errors, but simple Disconnection
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error reading message: %v", err)
			}
			break // Break the loop to close conn & Cleanup
		}

		var event Event
		if err = json.Unmarshal(data, &event); err != nil {
			log.Printf("error unmarshaling json: %v", err)
			continue
		}

		// Check if event in client to server
		if event.Type > 199 {
			continue
		}

		parsedPayload, err := parsePayload(event)
		if err != nil {
			fmt.Println(err)
			continue
		}

		event = Event{
			Type:    event.Type,
			Payload: parsedPayload,
		}

		if err := cm.routeEvent(event, c); err != nil {
			log.Printf("error routing event: %v", err)
		}

	}

}

// writeMessages is a process that listens for new messages to output to the Client.
//
// It will only send the client messages that are of event type server to client (>= 200)
func (c *Client) writeMessages(cm *ConnectionManager) {
	defer func() {

		// If client has existing session, remove it
		if _, ok := cm.sessions[c]; ok {
			cm.removeSession(c)
		}
		cm.removeClient(c)
	}()

	for {

		message, ok := <-c.egress
		// Ok will be false Incase the egress channel is closed
		if !ok {
			// Manager has closed this connection channel, so communicate that to frontend
			if err := c.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
				// Log that the connection is closed and the reason
				log.Println("connection closed: ", err)
			}
			// Return to close the goroutine
			return
		}

		// Check if event if server to client
		if message.Type < 200 {
			continue
		}

		// Write a Regular text message to the connection
		data, _ := json.Marshal(message)
		if err := c.connection.WriteMessage(websocket.TextMessage, data); err != nil {
			log.Println(err)
		}
		log.Println("sent message")

	}
}

func parsePayload(event Event) (interface{}, error) {
	switch event.Type {
	case EventSendMessage:
		payload, ok := event.Payload.(string)
		if !ok {
			return nil, fmt.Errorf("parsePaload: %v is of type %v not string", event.Payload, reflect.TypeOf(event.Payload))
		}
		return payload, nil
	case EventReqMatch:
		payload, ok := event.Payload.([]interface{})
		if !ok {
			return nil, fmt.Errorf("parsePaload: %v is of type %v not []interface{}", event.Payload, reflect.TypeOf(event.Payload))
		}
		payloadString := make([]string, len(payload))
		for i, v := range payload {
			str, ok := v.(string)
			if !ok {
				return nil, fmt.Errorf("type assertion failed for element: %v", v)
			}
			payloadString[i] = str
		}
		return payloadString, nil
	case EventReqSessionEnd:
		payload, ok := event.Payload.(struct{})
		if !ok {
			return nil, fmt.Errorf("parsePaload: %v is of type %v not struct", event.Payload, reflect.TypeOf(event.Payload))
		}
		return payload, nil

	default:
		fmt.Println("Unknown event type")
	}
	return nil, fmt.Errorf("parsePaload: %v is an invalid event type", event.Type)

}
