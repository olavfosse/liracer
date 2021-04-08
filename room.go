package main

import (
	"log"
	"sync"
)

type room struct {
	playersMu sync.Mutex
	// players is the set of players currently in Room.
	players map[*player]struct{}
	snippet string
}

// newRoom creates a new room with a random snippet.
func newRoom() *room {
	return &room{
		players: make(map[*player]struct{}),
		snippet: randomSnippet(),
	}
}

// register registers p to r. register is concurrency-safe.
func (r *room) register(p *player) {
	r.playersMu.Lock()
	defer r.playersMu.Unlock()
	r.players[p] = struct{}{}
	log.Printf("registered %v to room, there are now %d players in the room\n", p, len(r.players))
}

// unregister unregisters p from r. unregister is concurrency-safe.
func (r *room) unregister(p *player) {
	r.playersMu.Lock()
	defer r.playersMu.Unlock()
	delete(r.players, p)
	log.Printf("unregistered %v from room, there are now %d players in the room\n", p, len(r.players))
}

// sendToAllExcept sends bs to all players in r except p. sendToAllExcept is
// concurrency-safe.
func (r *room) sendToAllExcept(p *player, bs []byte) {
	r.playersMu.Lock()
	defer r.playersMu.Unlock()
	for pp := range r.players {
		if p != pp {
			pp.writeMessage(bs)
		}
	}
}
