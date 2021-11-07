package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"play.liracer.org/room"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// newWsHandler returns a http handler function which is used to set up the
// WebSocket endpoint that players interact with.
func newWsHandler() (func(http.ResponseWriter, *http.Request), error) {
	toRoomQueue, err := room.Start()
	if err != nil {
		return nil, err
	}

	nextPlayerID := room.PlayerID(1)
	nextPlayerIDMu := sync.Mutex{}

	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		// TODO: set read deadlines
		// TODO: set write deadline
		// TODO: remove incomingMsg.go
		// TODO: remove outgoingMsg.go
		nextPlayerIDMu.Lock()
		id := nextPlayerID
		nextPlayerID++
		nextPlayerIDMu.Unlock()

		toPlayerQueue := make(chan room.PlayerMessage, 1000) // read only
		go func() {
			for message := range toPlayerQueue {
				switch m := message.(type) {
				case room.ChatMessage_PlayerMessage:
					err := conn.WriteJSON(
						outgoingMsg{
							ChatMessageMsg: &ChatMessageOutgoingMsg{
								Sender:  m.Sender,
								Content: m.Content,
							},
						},
					)
					if err != nil {
						log.Printf("room: write to %d failed: %s\n", id, err)
						toRoomQueue <- room.Leave_RoomMessage(id)
						return
					}
				case room.NewRound_PlayerMessage:
					err := conn.WriteJSON(
						outgoingMsg{
							NewRoundMsg: &NewRoundOutgoingMsg{
								Snippet:    m.Snippet,
								NewRoundId: m.NewRoundID,
								// TODO: remove NewRoundID from client/frontend
							},
						},
					)
					if err != nil {
						log.Printf("room: write to %d failed: %s\n", id, err)
						toRoomQueue <- room.Leave_RoomMessage(id)
						return
					}
				case room.TypedCorrectChars_PlayerMessage:
					err := conn.WriteJSON(
						outgoingMsg{
							OpponentCorrectCharsMsg: &OpponentCorrectCharsOutgoingMsg{
								OpponentID:   m.PlayerID,
								CorrectChars: m.Chars,
							},
						},
					)
					if err != nil {
						log.Printf("room: write to %d failed: %s\n", id, err)
						toRoomQueue <- room.Leave_RoomMessage(id)
						return
					}
				}
			}
		}()

		toRoomQueue <- room.Join_RoomMessage{
			PlayerID:           id,
			PlayerMessageQueue: toPlayerQueue,
		}
		defer func() {
			toRoomQueue <- room.Leave_RoomMessage(id)
		}()

		for {
			_, bs, err := conn.ReadMessage()
			if err != nil {
				log.Printf("player %d: read failed: %s\n", id, err)
				return
			}
			log.Printf("player %d: read message %q\n", id, bs)
			var m incomingMsg
			err = json.Unmarshal(bs, &m)
			if err != nil {
				log.Printf("player %d: unmarshal failed: %s", id, err)
				continue
			}

			messageUnhandled := true

			if m.CorrectCharsMsg != nil {
				messageUnhandled = false
				toRoomQueue <- room.TypedCorrectChars_RoomMessage{
					PlayerID: id,
					RoundID:  m.CorrectCharsMsg.RoundId,
					Chars:    m.CorrectCharsMsg.CorrectChars,
				}
			}

			if m.ChatMessageMsg != nil {
				messageUnhandled = false
				toRoomQueue <- room.ChatMessage_RoomMessage{
					PlayerId: id,
					Content:  m.ChatMessageMsg.Content,
				}
			}
			if messageUnhandled {
				log.Printf("%d: unhandled message %q\n", id, bs)
			}
		}
	}, nil
}
