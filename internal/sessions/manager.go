package sessions

// import "github.com/majorbruteforce/hi-five/internal/broadcast"

// type Client = broadcast.Client

// // SessionManager manages chat sessions.
// type SessionManager struct {
// 	peers map[*Client]*Client
// }

// // NewSessionManager returns a new instance of a SessionManager
// func NewSessionManager() *SessionManager {
// 	return &SessionManager{
// 		peers: make(map[*Client]*Client),
// 	}
// }

// // createSession binds two existing connections to form a session.
// //
// // If a connection is already in session, an error is returned.
// func (sm *SessionManager) createSession(c1, c2 *Client) error {
// 	if _, ok := sm.peers[c1]; ok {
// 		// return fmt.Errorf("%s is already in session", c1.Connection.LocalAddr())
// 	}
// 	if _, ok := sm.peers[c2]; ok {
// 		// return fmt.Errorf("%s is already in session", c2.Connection.LocalAddr())
// 	}

// 	sm.peers[c1] = c2
// 	sm.peers[c2] = c1
// 	return nil
// }

// // removeSession unbinds two connections in an existing session.
// //
// // Either connection may be provided to refer to the session.
// func (sm *SessionManager) removeSession(c *Client) error {
// 	if _, ok := sm.peers[c]; ok {
// 		// return fmt.Errorf("%s has no existing sessions", c.Connection.LocalAddr())
// 	}

// 	// Other client in the session
// 	oc := sm.peers[c]

// 	delete(sm.peers, c)
// 	delete(sm.peers, oc)

// 	return nil
// }
