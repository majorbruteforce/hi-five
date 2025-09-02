package broadcast

import (
	"fmt"
	"log"
)

// createSession creates a new session if both clients are
// not in session.
//
// A EventNewSession is sent to both client's egress
func (cm *ConnectionManager) createSession(c1, c2 *Client) {
	cm.Lock()
	defer cm.Unlock()

	if _, ok := cm.sessions[c1]; ok {
		log.Printf("%v already in session", c1.connection.RemoteAddr())
		return
	}
	if _, ok := cm.sessions[c2]; ok {
		log.Printf("%v already in session", c2.connection.RemoteAddr())
		return
	}

	cm.sessions[c1] = c2
	cm.sessions[c2] = c1

	c1.egress <- Event{
		Type:    EventNewSession,
		Payload: c2.Profile,
	}

	c2.egress <- Event{
		Type:    EventNewSession,
		Payload: c1.Profile,
	}

}

// removeSession deletes an existing session.
//
// Either client may be passed as argument.
func (cm *ConnectionManager) removeSession(c *Client) {
	cm.Lock()
	defer cm.Unlock()

	if _, ok := cm.sessions[c]; !ok {
		log.Printf("client is not in session")
		return
	}

	// Other client in the session
	oc := cm.sessions[c]

	delete(cm.sessions, c)
	delete(cm.sessions, oc)

	oc.egress <- Event{
		Type:    EventSessionEnded,
		Payload: struct{}{},
	}
}

func (cm *ConnectionManager) CreateRandomSession() {
	for {
		if len(cm.clients) == 2 {
			break
		}
	}
	clients := make([]*Client, 0)
	for key := range cm.clients {
		clients = append(clients, key)
	}

	cm.createSession(clients[0], clients[1])
}

func (cm *ConnectionManager) CreateBatchSessions(in <-chan [][2]string) {
	for {
		matches, ok := <-in
		clientIDMap := make(map[string]*Client)
		for c := range cm.clients {
			clientIDMap[c.ID] = c
		}
		if !ok {
			fmt.Printf("session ingress channel is closed and drained")
			return
		}
		for _, m := range matches {
			cm.createSession(clientIDMap[m[0]], clientIDMap[m[1]])
		}
	}
}
