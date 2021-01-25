// Serve ./public
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("public")))
	address := "localhost:3000"
	fmt.Printf("front\t%s\n", address)
	err := http.ListenAndServe(address, nil)
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
