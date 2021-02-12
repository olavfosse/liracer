package main

import (
	"github.com/gorilla/websocket"
	"log"
)

type game struct {
	// players keeps track of the registered players. Each key in players
	// corresponds to a player.
	players map[*websocket.Conn]struct{}
	// register is used to register a player to the game.
	register chan *websocket.Conn
	// unregister is used to unregister a player from the game.
	unregister chan *websocket.Conn
}

func newGame() *game {
	return &game{
		players:    make(map[*websocket.Conn]struct{}),
		register:   make(chan *websocket.Conn),
		unregister: make(chan *websocket.Conn),
	}
}

func (g *game) run() {
	for {
		select {
		case player := <-g.register:
			g.players[player] = struct{}{}
			log.Println("registered player")
		case player := <-g.unregister:
			delete(g.players, player)
			log.Println("unregistered player")
		}
		log.Println("there are", len(g.players), "registered players")
	}
}

// writeJSONToAllExcept writes a JSON encoding of obj to all players except
// conn.
func (g *game) writeJSONToAllExcept(conn *websocket.Conn, obj interface{}) {
	for c := range g.players {
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
