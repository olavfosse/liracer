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
func newWsHandler() (func(http.ResponseWriter, *http.Request), error) {
	rm, err := newRoom()
	if err != nil {
		return nil, err
	}
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
				log.Printf("%v: read failed: %s\n", p, err)
				rm.handlePlayerLeft(p)
				return
			}
			log.Printf("%v: read message %q\n", p, bs)
			var m incomingMsg
			err = json.Unmarshal(bs, &m)
			if err != nil {
				log.Printf("%v: unmarshal failed: %s", p, err)
				continue
			}

			messageUnhandled := true

			if m.CorrectCharsMsg != nil {
				messageUnhandled = false
				rm.handlePlayerTypedCorrectChars(p, m.CorrectCharsMsg.CorrectChars)
			}

			if m.ChatMessageMsg != nil {
				messageUnhandled = false
				rm.handlePlayerSentChatMessage(p, m.ChatMessageMsg.Content)
			}
			if messageUnhandled {
				log.Printf("%v: unhandled message %q\n", p, bs)
			}
		}
	}, nil
}
