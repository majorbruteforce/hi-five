package broadcast

import "log"

func (cm *ConnectionManager) createSession(c1, c2 *Client) {
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
}
