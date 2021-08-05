package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	address, ok := os.LookupEnv("ADDRESS")
	if !ok {
		log.Fatalln(`environment variable "ADDRESS" not pressent`)
	}
	http.Handle("/", http.FileServer(http.Dir("public")))
	log.Println("listening on ", address)
	log.Fatalln(http.ListenAndServe(address, nil))
}
