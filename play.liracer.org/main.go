package main

import (
	"log"
	"net/http"
	"os"

	"github.com/caddyserver/certmagic"
)

func main() {
	// For publicFS to be defined, use `embed` or `donotembed` build tag.
	http.Handle("/", http.FileServer(http.FS(publicFS)))
	// NOTE: The game URLs have to be of a form something like
	//       https://play.liracer.org/?gameid=myepicgameid or
	//       https://play.liracer.org/id/anotherepicgameid since the
	//       previous URL form conflicts with the ws endpoint URL.
	handler, err := newWsHandler()
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/ws", handler)

	address, ok := os.LookupEnv("ADDRESS")
	if !ok {
		log.Fatalln(`environment variable "ADDRESS" not present`)
	}
	_, useHTTPS := os.LookupEnv("USE_HTTPS")
	if useHTTPS {
		log.Println("Listening on https://" + address)
		log.Fatalln(certmagic.HTTPS([]string{address}, nil))
	} else {
		log.Println("Listening on http://" + address)
		log.Fatalln(http.ListenAndServe(address, nil))
	}
}
