package broadcaster

type EventHandler func(event Event, c *Client) error

// setupEventHandlers configures and adds all handlers
func (cm *ConnectionManager) setupEventHandlers() {
	cm.handlers[EventSendMessage] = func(e Event, c *Client) error {

	}
}
