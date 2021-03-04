package game

import (
	"log"
	"sync"

	"github.com/fossegrim/play.liracer.org/snippet"
)

// player is an interface that allows us to call methods on the concrete
// type player.Player without importing the player package. Importing the player
// package would cause an import cycle, since the player package already imports
// the game package.
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

type Game struct {
	playersMu sync.Mutex
	// players is the set of players currently in Game.
	players map[player]struct{}
	snippet string
}

// New creates a new game with a random snippet.
func New() *Game {
	return &Game{
		players: make(map[player]struct{}),
		snippet: snippet.Random(),
	}
}

// Register registers p to g. Register is concurrency-safe.
func (g *Game) Register(p player) {
	g.playersMu.Lock()
	defer g.playersMu.Unlock()
	g.players[p] = struct{}{}
	log.Printf("Registered %v to game, there are now %d players in the game\n", p, len(g.players))
}

// Unregister unregisters p from g. Unregister is concurrency-safe.
func (g *Game) Unregister(p player) {
	g.playersMu.Lock()
	defer g.playersMu.Unlock()
	delete(g.players, p)
	log.Printf("Unregistered %v from game, there are now %d players in the game\n", p, len(g.players))
}

// SendToAllExcept sends bs to all players in g except p. SendToAllExcept is
// concurrency-safe.
func (g *Game) SendToAllExcept(p player, bs []byte) {
	g.playersMu.Lock()
	defer g.playersMu.Unlock()
	for pp := range g.players {
		if p != pp {
			pp.WriteMessage(bs)
		}
	}
}

// Snippet returns g's current snippet. Snippet is concurrency-safe.
func (g *Game) Snippet() string {
	return g.snippet
}
