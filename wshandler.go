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
		rm.handlePlayerJoined(p)

		for {
			_, bs, err := p.ReadMessage()
			if err != nil {
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

			if m.CorrectCharsMsg == nil {
				log.Printf("error: unhandled message: %q\n", bs)
				continue
			}
			rm.handlePlayerTypedCorrectChars(p, m.CorrectCharsMsg.CorrectChars)
		}
	}
}
