package sockets

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

func (sm *SocketManager) SendTo(userID string, msg []byte) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	if client, ok := sm.clients[userID]; ok {
		select {
		case client.Send <- msg:
		default:
		}
	}
}

func (sm *SocketManager) Broadcast(msg []byte) {
	sm.broadcast <- msg
}

func (c *Client) readPump(sm *SocketManager) {
	defer func() {
		sm.unregister <- c
	}()

	c.Conn.SetReadLimit(512)
	c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	c.Conn.SetPongHandler(func(string) error {
		c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}

		fmt.Printf("%s", msg)
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(30 * time.Second)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case msg, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.Conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				return
			}

		case <-ticker.C:
			c.Conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
