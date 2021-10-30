// +build embed
// also see public_donotembed.go

package main

import (
	"embed"
)

//go:embed public
var publicFS embed.FS
