package sockets

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"

	"github.com/majorbruteforce/hifive/internal/config"
	log "github.com/majorbruteforce/hifive/pkg/logger"
)

type Client struct {
	UserId string
	Conn   *websocket.Conn
	Send   chan []byte
}

type SocketManager struct {
	clients    map[string]*Client
	mu         sync.RWMutex
	upgrader   websocket.Upgrader
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
	cfg        config.Config
}

func NewSocketManager(cfg config.Config) *SocketManager {
	return &SocketManager{
		clients: make(map[string]*Client),
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin:     func(r *http.Request) bool { return true },
		},
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
		cfg:        cfg,
	}
}

func (sm *SocketManager) Run() {
	for {
		select {
		case client := <-sm.register:
			sm.mu.Lock()
			sm.clients[client.UserId] = client
			sm.mu.Unlock()
			go client.writePump()
			go client.readPump(sm)

		case client := <-sm.unregister:
			sm.mu.Lock()
			if _, ok := sm.clients[client.UserId]; ok {
				delete(sm.clients, client.UserId)
				close(client.Send)
				client.Conn.Close()
			}
			sm.mu.Unlock()

		case msg := <-sm.broadcast:
			sm.mu.RLock()
			for _, client := range sm.clients {
				select {
				case client.Send <- msg:
				default:
				}
			}
			sm.mu.RUnlock()
		}
	}
}

func (sm *SocketManager) HandleWS(w http.ResponseWriter, r *http.Request, userID string) {
	conn, err := sm.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error("upgrade:", err)
		return
	}

	client := &Client{
		UserId: userID,
		Conn:   conn,
		Send:   make(chan []byte, 256),
	}
	sm.register <- client
}

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
