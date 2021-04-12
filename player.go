package main

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type player struct {
	*websocket.Conn
	id id
}

type id int

// newPlayer creates a new player with a unique ID. newPlayer is concurrency
// safe.
func newPlayer(conn *websocket.Conn) *player {
	nextIDMu.Lock()
	defer nextIDMu.Unlock()
	p := &player{
		conn,
		nextID,
	}
	nextID++
	return p
}

var (
	nextIDMu sync.Mutex
	// nextID is the ID that the next created player will have.
	nextID id = 1
)

func (p *player) String() string {
	return fmt.Sprintf("player %d", p.id)
}
