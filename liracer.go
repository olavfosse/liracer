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
	id := singletonGame.RegisterPlayer(conn)
	for {
		var ccmfp correctCharsMessageFromPlayer
		if err := conn.ReadJSON(&ccmfp); err != nil {
			log.Println(err)
			singletonGame.Unregister <- id
			return
		}
		var ccmtp correctCharsMessageToPlayer
		ccmtp.correctCharsMessageFromPlayer = ccmfp
		ccmtp.PlayerId = id
		// NB: possible race condition
		singletonGame.WriteJSONToAllExcept(id, ccmtp)
	}
}

type baseMessage struct {
	MessageType string
}

type correctCharsMessageFromPlayer struct {
	baseMessage
	CorrectChars int
}

type correctCharsMessageToPlayer struct {
	correctCharsMessageFromPlayer
	PlayerId int
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
