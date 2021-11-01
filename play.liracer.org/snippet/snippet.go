package snippet

import (
	"embed"
	"fmt"
	"io/fs"
	"math/rand"
	"path/filepath"
	"strings"
)

//go:embed c/* go/* javascript/* shellscript/*
var snippetsFS embed.FS

type Snippet struct {
	Name     string
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
		code := string(bs)
		code = strings.Replace(code, "\r\n", "\n", -1)

		snip := Snippet{
			Name:     filepath.Base(path),
			Code:     code,
			Language: filepath.Base(filepath.Dir(path)),
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
	return nil
}
