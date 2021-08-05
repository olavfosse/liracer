package main

import (
	"log"
	"net/http"
	"os"

	"github.com/caddyserver/certmagic"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("public")))
	// NOTE: The game URLs have to be of a form something like
	//       https://play.liracer.org/?gameid=myepicgameid or
	//       https://play.liracer.org/id/anotherepicgameid since the
	//       previous URL form conflicts with the ws endpoint URL.
	http.HandleFunc("/ws", newWsHandler())

	address, ok := os.LookupEnv("ADDRESS")
	if !ok {
		log.Fatalln(`environment variable "ADDRESS" not pressent`)
	}
	_, useHTTPS := os.LookupEnv("USE_HTTPS")
	log.Println("listening on", address)
	if useHTTPS {
		log.Fatalln(certmagic.HTTPS([]string{address}, nil))
	} else {
		log.Fatalln(http.ListenAndServe(address, nil))
	}
}
