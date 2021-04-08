package main

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type player struct {
	id id

	connWriteMu sync.Mutex
	conn        *websocket.Conn
}

type id int

// newPlayer creates a new player with a guaranteed unique ID. newPlayer is
// concurrency safe.
func newPlayer(conn *websocket.Conn) *player {
	nextIDMu.Lock()
	defer nextIDMu.Unlock()
	p := &player{
		id:   nextID,
		conn: conn,
	}
	nextID++
	return p
}

var (
	nextIDMu sync.Mutex
	// nextID is the ID that the next created player will have.
	nextID id = 1
)

// writeMessage writes bs to p's underlying connection in a concurrency-safe way.
// It returns an error if the write failed.
func (p *player) writeMessage(bs []byte) (err error) {
	p.connWriteMu.Lock()
	defer p.connWriteMu.Unlock()

	return p.conn.WriteMessage(websocket.TextMessage, bs)
}

func (p *player) String() string {
	return fmt.Sprintf("Player %d", p.id)
}
