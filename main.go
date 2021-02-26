package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

const snippet = `package main

import fmt

func main() {
	fmt.Println("hello, world!")
}
`

func wsHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	p := newPlayer()

	// To prevent reading/writing concurrently to conn we discard it and
	// discard of the read/write methods outside the goroutine where they
	// should be used.
	writeMessage := conn.WriteMessage
	readMessage := conn.ReadMessage
	conn = nil
	go func() {
		wm := writeMessage
		writeMessage = nil // we may only write from the goroutine above

		for {
			bs := <-p.send
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
		// NOTE: Instead of manually iterating through the struct fields
		//       I could maybe use some form of reflection for this.
		if m.JoinGameIncomingMsg != nil {
			isMessageHandled = true
			bs, err := json.Marshal(outgoingMsg{
				GameId: gameId("dummyvalue"),
				GameStateOutgoingMsg: &GameStateOutgoingMsg{
					RoundId: 1, // dummy value
					Snippet: snippet,
				},
			})
			if err != nil {
				log.Println("error:", err)
			} else {
				p.send <- bs
			}
		}
		if !isMessageHandled {
			log.Printf("error: unhandled message: %q\n", bs)
		}
	}
}

//go:embed public
var embedee embed.FS

func main() {
	public, err := fs.Sub(embedee, "public")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	http.Handle("/", http.FileServer(http.FS(public)))
	// NOTE: The game URLs have to be of a form something like
	//       https://play.liracer.org/?gameid=myepicgameid or
	//       https://play.liracer.org/id/anotherepicgameid since the
	//       previous URL form conflicts with the ws endpoint URL.
	http.HandleFunc("/ws", wsHandler)
	address := "localhost:3000"
	log.Println("listening on", address)
	err = http.ListenAndServe(address, nil)
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
