package broadcast

import (
	"fmt"

	"github.com/majorbruteforce/hi-five/internal/matchmaker"
)

type EventHandler func(event Event, c *Client) error

// setupEventHandlers configures and adds all handlers
func (cm *ConnectionManager) setupEventHandlers(out chan<- matchmaker.Candidate) {
	cm.handlers[EventReqSendMessage] = func(e Event, c *Client) error {
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

	cm.handlers[EventReqMatch] = func(e Event, c *Client) error {
		out <- matchmaker.Candidate{ID: c.ID, Keywords: e.Payload.([]string)}
		c.egress <- Event{
			Type:    EventMatchmakingInProgress,
			Payload: struct{}{},
		}
		return nil
	}

	cm.handlers[EventReqSessionEnd] = func(event Event, c *Client) error {
		// TODO: check for existence of client and session
		// and return appropraite error
		cm.removeSession(c)
		return nil
	}

	cm.handlers[EventReqClientsStatus] = func(event Event, c *Client) error {

		cm.RLock()
		defer cm.RUnlock()

		o := len(cm.clients)
		m := len(cm.sessions) / 2
		payload := struct {
			OnlineCount int `json:"onlineCount"`
			MatchCount  int `json:"matchCount"`
		}{
			OnlineCount: o,
			MatchCount:  m,
		}
		c.egress <- Event{
			Type:    EventClientsStatus,
			Payload: payload,
		}

		return nil
	}
}
