package broadcast

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	connection *websocket.Conn
	egress     chan Event
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		connection: conn,
		egress:     make(chan Event),
	}
}

func (c *Client) readMessages(cm *ConnectionManager) {
	defer func() {
		c.connection.Close()
		cm.removeClient(c)
		// remove sessions if any
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
		err = json.Unmarshal(data, &event)

	}

}

// writeMessages is a process that listens for new messages to output to the Client
func (c *Client) writeMessages(cm *ConnectionManager) {
	defer func() {
		// Graceful close if this triggers a closing
		c.connection.Close()
		cm.removeClient(c)
		// @TODO: remove sessions if any
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
