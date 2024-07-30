package broadcast

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	// websocketUpgrader is used to upgrade incomming HTTP
	// requests into a persitent websocket connection
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

// ConnectionManager manages all websocket connections
type ConnectionManager struct {
	clients  map[*Client]struct{}
	handlers map[int]EventHandler
	sessions map[*Client]*Client
	sync.RWMutex
}

// NewConnectionManager returns a new instance of a ConnectionManager
func NewConnetionManager() *ConnectionManager {
	cm := &ConnectionManager{
		clients:  make(map[*Client]struct{}),
		handlers: make(map[int]EventHandler),
		sessions: make(map[*Client]*Client),
	}

	cm.setupEventHandlers()
	return cm

}

// addClient adds a new connection to clients list
func (cm *ConnectionManager) addClient(c *Client) {
	cm.Lock()
	defer cm.Unlock()

	cm.clients[c] = struct{}{}
}

// removeClient removes an existing client form the clients list
// and closes the connection
func (cm *ConnectionManager) removeClient(c *Client) {
	cm.Lock()
	defer cm.Unlock()

	if _, ok := cm.clients[c]; ok {
		c.connection.Close()
		delete(cm.clients, c)
		log.Printf("connection with %v closed", c.connection.RemoteAddr())
	}
}

// routeEvent is used to make sure the correct event goes into the correct handler
func (cm *ConnectionManager) routeEvent(event Event, c *Client) error {
	// Check if Handler is present in Map
	if handler, ok := cm.handlers[event.Type]; ok {
		// Execute the handler and return any err
		if err := handler(event, c); err != nil {
			return err
		}
		return nil
	} else {
		return ErrEventNotSupported
	}
}

// ServeConnections is a HTTP handler that accepts requests to create new
// web socket connections.
func (cm *ConnectionManager) ServeConnections(w http.ResponseWriter, r *http.Request) {
	// Begin by upgrading the HTTP request
	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("%s connected\n", conn.RemoteAddr())

	// Create New Client
	client := NewClient(conn)
	// Add the newly created client to the manager
	cm.addClient(client)
	// Start the read / write processes
	go client.readMessages(cm)
	go client.writeMessages(cm)
}

func (cm *ConnectionManager) Debug(w http.ResponseWriter, r *http.Request) {
	cm.RLock()
	defer cm.RUnlock()

	// Prepare the debug information
	debugInfo := struct {
		ClientCount  int               `json:"client_count"`
		Clients      []*Client         `json:"clients"`
		SessionCount int               `json:"session_count"`
		Sessions     map[string]string `json:"sessions"`
	}{
		ClientCount:  len(cm.clients),
		Clients:      make([]*Client, 0, len(cm.clients)),
		SessionCount: len(cm.sessions),
		Sessions:     make(map[string]string),
	}

	// Collect clients
	for client := range cm.clients {
		debugInfo.Clients = append(debugInfo.Clients, client)
	}

	// Collect session information
	for client, sessionClient := range cm.sessions {
		debugInfo.Sessions[client.ID] = sessionClient.ID
	}

	// Set response header and encode the debug information as JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(debugInfo); err != nil {
		http.Error(w, "Failed to encode debug information", http.StatusInternalServerError)
	}
}
