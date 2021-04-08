package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

func main() {
	public, err := fs.Sub(embedee, "public")
	if err != nil {
		log.Fatalln(err)
	}
	http.Handle("/", http.FileServer(http.FS(public)))
	// NOTE: The game URLs have to be of a form something like
	//       https://play.liracer.org/?gameid=myepicgameid or
	//       https://play.liracer.org/id/anotherepicgameid since the
	//       previous URL form conflicts with the ws endpoint URL.
	http.HandleFunc("/ws", wsHandler)
	address := "localhost:3000"
	log.Println("listening on", address)
	log.Fatalln(http.ListenAndServe(address, nil))
}

//go:embed public
var embedee embed.FS
