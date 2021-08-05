package main

import (
	"log"
	"net/http"
	"os"
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
	log.Println("listening on", address)
	log.Fatalln(http.ListenAndServe(address, nil))
}
