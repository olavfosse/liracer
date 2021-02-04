// Serve ./public
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"unicode/utf8"

	"github.com/gorilla/websocket"
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

	conn.WriteMessage(websocket.TextMessage, []byte("ping"))
	log.Println("wrote: ping")

	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		if !utf8.Valid(p) {
			log.Println(p, "is not entirely valid UTF-8")
			return
		}
		s := string(p)
		log.Println("read:", s)

		if s == "ping" {
			conn.WriteMessage(websocket.TextMessage, []byte("pong"))
			log.Println("wrote: pong")
		}
	}
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
