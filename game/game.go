package game

import (
	"github.com/gorilla/websocket"
	"log"
)

type Game struct {
	// Players keeps track of the registered players. Each key in Players
	// corresponds to a player.
	Players map[*websocket.Conn]struct{}
	// Register is used to register a player to the game.
	Register chan *websocket.Conn
	// Unregister is used to unregister a player from the game.
	Unregister chan *websocket.Conn
}

func NewGame() *Game {
	return &Game{
		Players:    make(map[*websocket.Conn]struct{}),
		Register:   make(chan *websocket.Conn),
		Unregister: make(chan *websocket.Conn),
	}
}

func (g *Game) Run() {
	for {
		select {
		case player := <-g.Register:
			g.Players[player] = struct{}{}
			log.Println("registered player")
		case player := <-g.Unregister:
			delete(g.Players, player)
			log.Println("unregistered player")
		}
		log.Println("there are", len(g.Players), "registered Players")
	}
}

// WriteJSONToAllExcept writes a JSON encoding of obj to all Players except
// conn.
func (g *Game) WriteJSONToAllExcept(conn *websocket.Conn, obj interface{}) {
	for c := range g.Players {
		if c == conn {
			continue
		}
		// PERFORMANCE: this encodes obj as JSON once for each player which is
		// very inefficient.
		if err := c.WriteJSON(obj); err != nil {
			log.Println(err)
		}
	}
}
