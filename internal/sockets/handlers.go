package sockets

import (
	"encoding/json"

	"github.com/majorbruteforce/hifive/internal/events"
)

type EventHandler func(event events.Event, c *Client) error

func (sm *SocketManager) setupEventHandlers() {

	sm.handlers[events.EventConnectionSuccessful] = func(event events.Event, c *Client) error {
		msg, err := json.Marshal(event)
		if err != nil {
			return err
		}

		c.SendMsg(msg)
		return nil
	}

}
