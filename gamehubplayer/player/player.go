package player

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Player struct {
	ID ID

	connWriteMu sync.Mutex
	conn        *websocket.Conn
}

type ID int

// New creates a new player with a guaranteed unique ID. New is concurrency
// safe.
func New(conn *websocket.Conn) *Player {
	nextIDMu.Lock()
	defer nextIDMu.Unlock()
	p := &Player{
		ID:   nextID,
		conn: conn,
	}
	nextID++
	return p
}

var (
	// nextID is the ID that the next created player will have.
	nextID   ID = 1
	nextIDMu sync.Mutex
)

// WriteMessage writes bs to p's underlying connection in a concurrency-safe way.
// It returns an error if the write failed.
func (p *Player) WriteMessage(bs []byte) (err error) {
	p.connWriteMu.Lock()
	defer p.connWriteMu.Unlock()

	return p.conn.WriteMessage(websocket.TextMessage, bs)
}
