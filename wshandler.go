package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// TODO: remove global variable
var singletonRoom *room

func init() {
	singletonRoom = newRoom()
}

// wsHandler is a http handler function, for use with http.HandleFunc, which
// is used to set up the WebSocket endpoint that players interact with.
func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	p := newPlayer(conn)

	for {
		_, bs, err := p.conn.ReadMessage()
		if err != nil {
			log.Println("error(closing connection):", err)
			singletonRoom.unregister(p)
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
		if m.JoinRoomMsg != nil {
			isMessageHandled = true

			singletonRoom.register(p)
			bs, err := json.Marshal(
				outgoingMsg{
					SetRoomStateMsg: &SetRoomStateOutgoingMsg{
						Snippet: singletonRoom.snippet,
					},
				},
			)
			if err != nil {
				log.Println("error:", err)
				panic("marshalling a outgoingMsg should never result in an error")
			}
			err = p.writeMessage(bs)
			if err != nil {
				log.Println("error(closing connection):", err)
				singletonRoom.unregister(p)
				return
			}
			log.Printf("wrote: %q\n", bs)
		}
		if m.CorrectCharsMsg != nil {
			isMessageHandled = true

			bs, err := json.Marshal(
				outgoingMsg{
					OpponentCorrectCharsMsg: &OpponentCorrectCharsIncomingMsg{
						OpponentID:   p.id,
						CorrectChars: m.CorrectCharsMsg.CorrectChars,
					},
				},
			)
			if err != nil {
				log.Println("error:", err)
				panic("marshalling a outgoingMsg should never result in an error")
			}
			singletonRoom.sendToAllExcept(p, bs)
		}
		if !isMessageHandled {
			log.Printf("error: unhandled message: %q\n", bs)
		}
	}
}
