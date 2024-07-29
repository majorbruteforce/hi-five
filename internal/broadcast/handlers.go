package broadcast

import "fmt"

type EventHandler func(event Event, c *Client) error

// setupEventHandlers configures and adds all handlers
func (cm *ConnectionManager) setupEventHandlers() {
	cm.handlers[EventSendMessage] = func(e Event, c *Client) error {
		if _, ok := cm.sessions[c]; ok {
			msg := e.Payload
			cm.sessions[c].egress <- Event{
				Type:    EventNewMessage,
				Payload: msg,
			}
		} else {
			return fmt.Errorf("%s not in session, trying to send message", c.ID)
		}
		return nil
	}
}
