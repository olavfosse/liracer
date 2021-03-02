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
	p := New(conn)

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
