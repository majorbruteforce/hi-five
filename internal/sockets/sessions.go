package sockets

import "fmt"

func (sm *SocketManager) CreateSession(c1, c2 *Client) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	if _, ok := sm.sessions[c1]; ok {
		return fmt.Errorf("Client(id=%s) is already in session", c1.UserId)
	}

	if _, ok := sm.sessions[c2]; ok {
		return fmt.Errorf("Client(id=%s) is already in session", c2.UserId)
	}

	sm.sessions[c1] = c2
	sm.sessions[c2] = c1

	return nil
}