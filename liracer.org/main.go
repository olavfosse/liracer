package main

import (
	"log"
	"net/http"
	"os"

	"github.com/caddyserver/certmagic"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("public")))

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
