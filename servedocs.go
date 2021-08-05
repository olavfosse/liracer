package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("docs")))
	log.Println("listening on localhost:3210")
	log.Fatalln(http.ListenAndServe(":3210", nil))
}
