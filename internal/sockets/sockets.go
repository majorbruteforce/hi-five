package sockets

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"

	"github.com/majorbruteforce/hifive/internal/config"
	log "github.com/majorbruteforce/hifive/pkg/logger"
)

type SocketManager struct {
	clients    map[string]*Client
	handlers   map[int]EventHandler
	sessions   map[*Client]*Client
	mu         sync.RWMutex
	upgrader   websocket.Upgrader
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
	cfg        config.Config
}

func NewSocketManager(cfg config.Config) *SocketManager {
	sm := &SocketManager{
		clients:  make(map[string]*Client),
		handlers: make(map[int]EventHandler),
		sessions: make(map[*Client]*Client),
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

	sm.setupEventHandlers()
	return sm
}

func (sm *SocketManager) Run() {
	for {
		select {
		case client := <-sm.register:
			sm.handleRegistration(client)

		case client := <-sm.unregister:
			sm.handleUnregistration(client)

		case msg := <-sm.broadcast:
			sm.handleBroadcast(msg)
		}
	}
}

func (sm *SocketManager) HandleWS(w http.ResponseWriter, r *http.Request, userID string) {
	conn, err := sm.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error(err)
		return
	}

	client := &Client{
		UserId: userID,
		Conn:   conn,
		Send:   make(chan []byte, 256),
	}
	sm.register <- client
}

func (sm *SocketManager) handleRegistration(c *Client) {
	sm.mu.Lock()
	sm.clients[c.UserId] = c
	sm.mu.Unlock()
	go c.writePump()
	go c.readPump(sm)

	log.Log.Infof("Client(id=%s) registered", c.UserId)
}

func (sm *SocketManager) handleUnregistration(c *Client) {
	sm.mu.Lock()
	if _, ok := sm.clients[c.UserId]; ok {
		delete(sm.clients, c.UserId)
		close(c.Send)
		c.Conn.Close()
	}
	sm.mu.Unlock()

	log.Log.Infof("Client(id=%s) unregistered", c.UserId)
}

func (sm *SocketManager) handleBroadcast(msg []byte) {
	sm.mu.RLock()
	for _, client := range sm.clients {
		select {
		case client.Send <- msg:

		default:
		}
	}
	sm.mu.RUnlock()

	log.Log.Infof("Broadcast message sent")
}

func (sm *SocketManager) RegisterWSHandler() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("userID")
		if userID == "" {
			http.Error(w, "missing userID", http.StatusBadRequest)
			return
		}
		sm.HandleWS(w, r, userID)
	})
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
