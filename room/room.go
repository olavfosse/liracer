package room

import (
	"log"
	"sync"

	"github.com/fossegrim/play.liracer.org/snippet"
)

// player is an interface that allows us to call methods on the concrete
// type player.Player without importing the player package. Importing the player
// package would cause an import cycle, since the player package already imports
// the room package.
//
// Thanks to mrig on freenode for suggesting this method of avoiding an import
// cycle!
//
// The methods on the player interface are intentionally undocumented. Since
// there is only one concrete type intended to satisfy the interface, the
// method documentation lives there instead.
type player interface {
	WriteMessage([]byte) error
	String() string
}

type Room struct {
	playersMu sync.Mutex
	// players is the set of players currently in Room.
	players map[player]struct{}
	snippet string
}

// New creates a new room with a random snippet.
func New() *Room {
	return &Room{
		players: make(map[player]struct{}),
		snippet: snippet.Random(),
	}
}

// Register registers p to r. Register is concurrency-safe.
func (r *Room) Register(p player) {
	r.playersMu.Lock()
	defer r.playersMu.Unlock()
	r.players[p] = struct{}{}
	log.Printf("Registered %v to room, there are now %d players in the room\n", p, len(r.players))
}

// Unregister unregisters p from r. Unregister is concurrency-safe.
func (r *Room) Unregister(p player) {
	r.playersMu.Lock()
	defer r.playersMu.Unlock()
	delete(r.players, p)
	log.Printf("Unregistered %v from room, there are now %d players in the room\n", p, len(r.players))
}

// SendToAllExcept sends bs to all players in r except p. SendToAllExcept is
// concurrency-safe.
func (r *Room) SendToAllExcept(p player, bs []byte) {
	r.playersMu.Lock()
	defer r.playersMu.Unlock()
	for pp := range r.players {
		if p != pp {
			pp.WriteMessage(bs)
		}
	}
}

// Snippet returns r's current snippet. Snippet is concurrency-safe.
func (r *Room) Snippet() string {
	return r.snippet
}
