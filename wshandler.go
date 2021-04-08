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

// newWsHandler returns a http handler function which is used to set up the
// WebSocket endpoint that players interact with.
func newWsHandler() func(http.ResponseWriter, *http.Request) {
	rm := newRoom()
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		p := newPlayer(conn)

		rm.register(p)
		bs, err := json.Marshal(
			outgoingMsg{
				SetRoomStateMsg: &SetRoomStateOutgoingMsg{
					Snippet: rm.snippet,
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
			rm.unregister(p)
			return
		}
		log.Printf("wrote: %q\n", bs)

		for {
			_, bs, err := p.conn.ReadMessage()
			if err != nil {
				log.Println("error(closing connection):", err)
				rm.unregister(p)
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
				rm.sendToAllExcept(p, bs)
			}
			if !isMessageHandled {
				log.Printf("error: unhandled message: %q\n", bs)
			}
		}
	}
}
