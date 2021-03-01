package player

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fossegrim/play.liracer.org/snippet"
	"github.com/gorilla/websocket"
)

type PlayerID int

type Player struct {
	// ID identifies Player.
	ID PlayerID
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

// New creates a Player with a unique id. It initializes send, but does not
// start reading from it. That job is reserved for the caller.
func New() *Player {
	return &Player{
		ID:   <-playerIdEmitter,
		Send: make(chan []byte),
	}
}

// HandlerFunc is used with http.HandleFunc to set up a websocket endpoint.
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
	p := New()

	// To prevent reading/writing concurrently to conn we discard it and
	// discard of the read/write methods outside the goroutine where they
	// should be used.
	writeMessage := conn.WriteMessage
	readMessage := conn.ReadMessage
	conn = nil
	go func() {
		wm := writeMessage
		writeMessage = nil // we may only write from within this goroutine

		for {
			bs := <-p.Send
			err := wm(websocket.TextMessage, bs)
			log.Printf("wrote: %q\n", bs)
			if err != nil {
				// TODO: actually close it
				log.Println("error(closing connection):", err)
				return
			}
		}
	}()

	for {
		_, bs, err := readMessage()
		if err != nil {
			// TODO: actually close it
			log.Println("error(closing connection):", err)
			return
		}
		log.Printf("read: %q\n", bs)
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
