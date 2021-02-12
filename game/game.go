package game

import (
	"github.com/gorilla/websocket"
	"log"
)

// NB: The method of registering players and retrieving their ids /feels/ very
//     error prone.

type Game struct {
	// Players keeps track of the registered players.
	Players map[int]*websocket.Conn
	// register is used to register a player to the game.
	register chan *websocket.Conn
	// registeredIds publishes the ids of newly registered players.
	registeredIds chan int
	// Unregister is used to unregister a player from the game by id.
	Unregister chan int
}

func NewGame() *Game {
	return &Game{
		Players:       make(map[int]*websocket.Conn),
		register:      make(chan *websocket.Conn),
		registeredIds: make(chan int),
		Unregister:    make(chan int),
	}
}

func (g *Game) Run() {
	// nextId is the id that the next registered player will use.
	nextId := 0
	for {
		select {
		case conn := <-g.register:
			g.Players[nextId] = conn
			g.registeredIds <- nextId
			log.Println("registered player", nextId)
			nextId++
		case id := <-g.Unregister:
			delete(g.Players, id)
			log.Println("unregistered player", id)
		}
		log.Println("there are", len(g.Players), "registered Players")
	}
}

// WriteJSONToAllExcept writes a JSON encoding of obj to all Players except
// the player with id idToExclude.
func (g *Game) WriteJSONToAllExcept(idToExclude int, obj interface{}) {
	for id, conn := range g.Players {
		if id == idToExclude {
			continue
		}
		// PERFORMANCE: this encodes obj as JSON once for each player which is
		//              very inefficient.
		if err := conn.WriteJSON(obj); err != nil {
			log.Println(err)
		}
	}
}

// RegisterPlayer registers a player with conn conn and returns its id.
func (g *Game) RegisterPlayer(conn *websocket.Conn) int {
	go func() {
		g.register <- conn
	}()
	id := <-g.registeredIds
	return id
}
