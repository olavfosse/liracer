// +build embed
// also see public_donotembed.go

package main

import (
	"embed"
	"io/fs"
)

//go:embed public
var publicEmbedFS embed.FS

var publicFS fs.FS

func init() {
	var err error
	publicFS, err = fs.Sub(publicEmbedFS, "public")
	if err != nil {
		panic(err)
	}
}
