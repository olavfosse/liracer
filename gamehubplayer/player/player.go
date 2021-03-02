package player

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

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
		conn: conn,
	}

	for {
		_, bs, err := p.conn.ReadMessage()
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
				panic("marshalling an outgoingMsg should never result in an error")
			}
			err = p.WriteMessage(bs)
			if err != nil {
				// TODO: actually close it
				log.Println("error(closing connection):", err)
				return
			}
			log.Printf("wrote: %q\n", bs)
		}
		if !isMessageHandled {
			log.Printf("error: unhandled message: %q\n", bs)
		}
	}
}

// WriteMessage writes bs to p's underlying connection in a concurrency-safe way.
// It returns an error if the write failed.
func (p *Player) WriteMessage(bs []byte) (err error) {
	p.connWriteMu.Lock()
	defer p.connWriteMu.Unlock()

	return p.conn.WriteMessage(websocket.TextMessage, bs)
}

type PlayerID int

type Player struct {
	// ID identifies Player.
	ID PlayerID

	connWriteMu sync.Mutex
	conn        *websocket.Conn
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
