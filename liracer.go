// Serve ./public
package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"

	"github.com/fossegrim/play.liracer.org/game"
)

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
	singletonGame.Register <- conn

	for {
		var ccm correctCharsMessage
		if err := conn.ReadJSON(&ccm); err != nil {
			log.Println(err)
			singletonGame.Unregister <- conn
			return
		}
		// NB: possible race condition
		singletonGame.WriteJSONToAllExcept(conn, ccm)
	}
}

type baseMessage struct {
	MessageType string
}

type correctCharsMessage struct {
	baseMessage
	CorrectChars int
}

var singletonGame *game.Game

func init() {
	singletonGame = game.NewGame()
	go singletonGame.Run()
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("public")))
	// NOTE: The game URLs have to be of a form something like
	//       https://play.liracer.org/?gameid=myepicgameid or
	//       https://play.liracer.org/id/anotherepicgameid since the
	//       previous URL form conflicts with the ws endpoint URL.
	http.HandleFunc("/ws", wsHandler)
	address := "localhost:3000"
	log.Println("listening on", address)
	err := http.ListenAndServe(address, nil)
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
