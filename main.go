package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/fossegrim/play.liracer.org/gamehubplayer/player"
)

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
	http.HandleFunc("/ws", player.WsHandler)
	address := "localhost:3000"
	log.Println("listening on", address)
	err = http.ListenAndServe(address, nil)
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

//go:embed public
var embedee embed.FS
