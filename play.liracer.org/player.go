package main

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type player struct {
	*websocket.Conn
	id playerId
}

type playerId int

// newPlayer creates a new player with a unique ID. newPlayer is concurrency
// safe.
func newPlayer(conn *websocket.Conn) *player {
	nextPlayerIDMu.Lock()
	defer nextPlayerIDMu.Unlock()
	p := &player{
		conn,
		nextPlayerId,
	}
	nextPlayerId++
	return p
}

var (
	nextPlayerIDMu sync.Mutex
	// nextPlayerId is the id that the next created player will have.
	nextPlayerId playerId = 1
)

func (p *player) String() string {
	return fmt.Sprintf("player %d", p.id)
}
