package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("docs")))
	log.Println("Listening on http://localhost:3211")
	log.Fatalln(http.ListenAndServe(":3211", nil))
}
