package snippet

import (
	"embed"
	"fmt"
	"io/fs"
	"math/rand"
	"path/filepath"
	"strings"
	"unicode"
)

//go:embed c/* go/* javascript/* shellscript/*
var snippetsFS embed.FS

type Snippet struct {
	Name string
	// Code is ASCII-encoded, because I want a fixed-width encoding.
	// If we need non-ASCII characters, I might switch to UTF-16.
	Code     string
	Language string
}

type SnippetSet struct {
	snippetSet []Snippet
}

func ParseSnippetSet() (*SnippetSet, error) {
	snippetSet := &SnippetSet{
		[]Snippet{},
	}
	err := fs.WalkDir(snippetsFS, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		bs, err := fs.ReadFile(snippetsFS, path)
		if err != nil {
			return err
		}
		code := string(bs) // we validate that this is ascii later
		code = strings.Replace(code, "\r\n", "\n", -1)

		lang := filepath.Base(filepath.Dir(path))
		if lang == "go" {
			// To prevent Go snippets from being treated like part of the codebase, and therefore causing VSCode "Problems", we prefix them all with "//go:build ignore\n". We do not however wish "//go:build ignore\n" to be included in snippet code served to the users.
			// See also https://forum.golangbridge.org/t/how-to-ignore-files-in-vscode-go/25244.
			code = strings.Replace(code, "//go:build ignore\n", "", 1)
		}

		snip := Snippet{
			Name:     filepath.Base(path),
			Code:     code,
			Language: lang,
		}
		err = validate(snip)
		if err != nil {
			return err
		}
		snippetSet.snippetSet = append(snippetSet.snippetSet, snip)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return snippetSet, nil
}

// Random returns a random Snippet.
func (snippetSet *SnippetSet) Random() Snippet {
	return snippetSet.snippetSet[rand.Intn(len(snippetSet.snippetSet))]
}

// Get returns the first snippet s for which s.Name == name.  If there
// is no such snippet, nil is returned.
func (snippetSet *SnippetSet) Get(name string) *Snippet {
	for _, s := range snippetSet.snippetSet {
		if s.Name == name {
			return &s
		}
	}
	return nil
}

// validate checks if snip is valid, if it is not an error is logged and the
// program is terminated.
func validate(snip Snippet) error {
	if snip.Code[len(snip.Code)-1] != '\n' {
		return fmt.Errorf("invalid snippet %s/%s: last character is not newline", snip.Language, snip.Name)
	}

	if i := firstNonASCIIByteIndex(snip.Code); i != -1 {
		return fmt.Errorf(
			"invalid snippet %s/%s: byte at index %d is not ascii, bits = %08b",
			snip.Language,
			snip.Name,
			i,
			snip.Code[i],
		)
	}
	return nil
}

// If s is valid ASCII, returns -1, otherwise returns index of first non-ASCII byte.
func firstNonASCIIByteIndex(s string) int {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return i
		}
	}
	return -1
}
