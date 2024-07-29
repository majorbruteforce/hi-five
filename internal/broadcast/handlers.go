package broadcast

import "fmt"

type EventHandler func(event Event, c *Client) error

// setupEventHandlers configures and adds all handlers
func (cm *ConnectionManager) setupEventHandlers() {
	cm.handlers[EventSendMessage] = func(e Event, c *Client) error {
		msg := e.Payload
		// Send message to map of c
		fmt.Printf("%v: %s", c.connection.RemoteAddr(), msg)
		return nil
	}
}
