package snippet

import (
	"embed"
	"log"
	"math/rand"
)

//go:embed c/* go/* javascript/* shellscript/*
var snippetsFS embed.FS

var snippets []Snippet

func init() {
	languageDirs, err := snippetsFS.ReadDir(".")
	if err != nil {
		log.Fatalln("snippets:", err)
	}
	for _, languageDir := range languageDirs {
		snippetFiles, err := snippetsFS.ReadDir(languageDir.Name())
		if err != nil {
			log.Fatalln("snippets:", err)
		}
		for _, snippetFile := range snippetFiles {
			name := languageDir.Name() + "/" + snippetFile.Name()
			bs, err := snippetsFS.ReadFile(name)
			if err != nil {
				log.Fatalln("snippets:", err)
			}
			code := string(bs)
			snip := Snippet{
				Name:     snippetFile.Name(),
				Code:     code,
				Language: languageDir.Name(),
			}
			validate(snip)
			snippets = append(snippets, snip)
		}
	}
}

// validate checks if snip is valid, if it is not an error is logged and the
// program is terminated.
func validate(snip Snippet) {
	if snip.Code[len(snip.Code)-1] != '\n' {
		log.Fatalf("snippet: invalid snippet %s/%s: last character is not newline\n", snip.Language, snip.Name)
	}
}

// Random returns a random Snippet.
func Random() Snippet {
	return snippets[rand.Intn(len(snippets))]
}

type Snippet struct {
	Name     string
	Code     string
	Language string
}
