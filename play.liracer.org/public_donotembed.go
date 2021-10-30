// +build donotembed
// also see public_embed.go

package main

import "os"

var publicFS = os.DirFS("public/")
