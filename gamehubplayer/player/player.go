package player

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fossegrim/play.liracer.org/snippet"
	"github.com/gorilla/websocket"
)

// HandlerFunc is a http handler function, for use with http.HandleFunc, which
// is used to set up a WebSocket endpoint that players interact with.
func HandlerFunc(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	p := &Player{
		ID:   <-playerIdEmitter,
		conn: conn, // MUST NOT BE WRITTEN DIRECTLY TO.
		Send: make(chan []byte),
	}
	go p.pumpFromSend()
	p.handleIncomingMessages()
}

// pumpFromSend continually pumps the messages from p.Send to p's underlying
// connection.
//
// THIS IS THE ONLY FUNCTION PERMITTED TO WRITE DIRECTLY TO p's UNDERLYING
// CONNECTION. ONLY ONE pumpFromSend GOROUTINE IS PERMITTED TO EXIST PER p.
func (p *Player) pumpFromSend() {
	for {
		bs := <-p.Send
		err := p.conn.WriteMessage(websocket.TextMessage, bs)
		if err != nil {
			// TODO: actually close it
			log.Println("error(closing connection):", err)
			return
		}
	}
}

// handleIncomingMessages continually reads and handles incoming messages from
// p's underlying connection.
//
// THIS IS THE ONLY FUNCTION PERMITTED TO READ DIRECTLY FROM p's UNDERLYING
// CONNECTION. ONLY ONE handleIncomingMessages GOROUTINE IS PERMITTED TO EXIST
// PER p.
func (p *Player) handleIncomingMessages() {
	for {
		_, bs, err := p.conn.ReadMessage()
		if err != nil {
			// TODO: actually close it
			log.Println("error(closing connection):", err)
			return
		}
		var m incomingMsg
		err = json.Unmarshal(bs, &m)
		if err != nil {
			log.Println("error:", err)
			continue
		}

		isMessageHandled := false
		if m.JoinGameMsg != nil {
			isMessageHandled = true
			bs, err := json.Marshal(
				outgoingMsg{
					SetGameStateMsg: &SetGameStateOutgoingMsg{
						Snippet: snippet.Random(),
					},
				},
			)
			if err != nil {
				log.Println("error:", err)
			} else {
				p.Send <- bs
			}
		}
		if !isMessageHandled {
			log.Printf("error: unhandled message: %q\n", bs)
		}

	}
}

type PlayerID int

type Player struct {
	// ID identifies Player.
	ID PlayerID
	// conn is Player's underlying connection.
	conn *websocket.Conn
	// Send is a channel for writing to Player's underlying connection.
	Send chan []byte
}

// playerIdEmitter generates PlayerIds. It first produces PlayerId(1). Thereafter
// it always produces PlayerIds one greater than the previous, I.e PlayerId(1),
// PlayerId(2), ..., PlayerId(n).
var playerIdEmitter = make(chan PlayerID)

func init() {
	// No cleanup is necessary
	go func() {
		nextID := PlayerID(1)
		for {
			playerIdEmitter <- nextID
			nextID++
		}
	}()
}
