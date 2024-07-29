package broadcast

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"

	"github.com/google/uuid"
)

type Client struct {
	ID         string
	connection *websocket.Conn
	egress     chan Event
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		ID:         uuid.New().String(),
		connection: conn,
		egress:     make(chan Event),
	}
}

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
		if err := cm.routeEvent(event, c); err != nil {
			log.Printf("error routing event: %v", err)
		}

	}

}

// writeMessages is a process that listens for new messages to output to the Client
func (c *Client) writeMessages(cm *ConnectionManager) {
	defer func() {

		// If client has existing session, remove it
		if _, ok := cm.sessions[c]; ok {
			cm.removeSession(c)
		}
		cm.removeClient(c)
	}()

	for {
		select {
		case message, ok := <-c.egress:
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
			// Write a Regular text message to the connection
			data, _ := json.Marshal(message)
			if err := c.connection.WriteMessage(websocket.TextMessage, data); err != nil {
				log.Println(err)
			}
			log.Println("sent message")
		}

	}
}
