package player

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type IncomingMsg struct {
	JoinGameMsg     *JoinGameMsg
	CorrectCharsMsg *CorrectCharsMsg
}

type JoinGameMsg struct {
	GameId string
}

type CorrectCharsMsg struct {
	CorrectChars int
}

type OutboundMsg struct {
	SnippetMsg *SnippetMsg
}

type SnippetMsg struct {
	Snippet string
}

type Id int

type Player struct {
	// Id identifies Player.
	Id int
	// GameId identifies Player's current game. If gameId == "" Player is
	// not in a game.
	GameId string
	// Send is a channel for writing to Player's underlying connection.
	Send chan []byte
	// conn is the underlying connection of Player.
	conn *websocket.Conn
}

// Run handles the incoming requests to p's underlying websocket connection.
func (p *Player) Run() {
	go func() {
		// To prevent race conditions this loop is the only place where
		// we should write to p.conn.
		err := p.conn.WriteMessage(websocket.TextMessage, <-p.Send)
		if err != nil {
			log.Println("error(closing connection):", err)
			return
		}
	}()
	for {
		// To prevent race conditions this loop is the only place where
		// we should read from p.conn.
		var m IncomingMsg
		err := p.conn.ReadJSON(&m)
		if err != nil {
			log.Println("error(closing connection):", err)
			return
		}

		isMessageHandled := false
		// NOTE: Instead of manually iterating through the struct fields
		//       I could maybe use some form of reflection for this.
		if m.JoinGameMsg != nil {
			isMessageHandled = true
			s := `package main

import fmt

func main() {
	fmt.Println("hello, world!")
}
`
			bs, err := json.Marshal(OutboundMsg{&SnippetMsg{s}})
			if err != nil {
				log.Println("error:", err)
			} else {
				p.Send <- bs
			}
		}
		if m.CorrectCharsMsg != nil {
			isMessageHandled = true
			// Do nothing (temporarily)
		}
		if !isMessageHandled {
			log.Println("error: incoming message has neither a JoinGameMsg or CorrectCharsMsg field")
		}
	}
}

// A PlayerCreator provides a concurrency safe way of creating Player objects
// with unique ids, I.e different from any other Player created with the same
// PlayerCreator.
type PlayerCreator struct {
	// idEmitter emits unique ids.
	idEmitter chan int
}

// NewPlayer creates a Player with a unique id.
func (pc *PlayerCreator) NewPlayer(conn *websocket.Conn) *Player {
	p := &Player{
		Id:     <-pc.idEmitter,
		GameId: "", // initialy player is not in a game
		conn:   conn,
		Send:   make(chan []byte),
	}
	log.Println("created player with id", p.Id)
	return p
}

// NewPlayerCreator creates a ready-to-use PlayerCreator.
func NewPlayerCreator() *PlayerCreator {
	pc := PlayerCreator{
		idEmitter: make(chan int),
	}

	go func() {
		nextId := 1
		for {
			pc.idEmitter <- nextId
			nextId++
		}

	}()
	return &pc
}
