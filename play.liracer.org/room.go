package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"
	"unicode/utf8"

	"github.com/gorilla/websocket"
	"play.liracer.org/snippet"
)

type room struct {
	sync.Mutex
	// players is the set of players currently in room.
	players map[*player]struct{}
	// playerTypedFirstCorrectChar maps players to the points in time
	// when they typed their first char correctly.
	playerTypedFirstCorrectChar map[*player]time.Time
	snippet                     snippet.Snippet
	roundId                     roundId
}

type roundId int

// newRoom creates a new room with a random snippet.
func newRoom() *room {
	return &room{
		players:                     make(map[*player]struct{}),
		playerTypedFirstCorrectChar: make(map[*player]time.Time),
		snippet:                     snippet.Random(),
		roundId:                     1,
	}
}

// sendTo sends bs to p. If an error occurs p it is logged.  sendTo is
// not concurrency safe.
func (r *room) sendTo(p *player, bs []byte) {
	err := p.WriteMessage(websocket.TextMessage, bs)
	if err != nil {
		log.Printf("room: write to %v failed: %s\n", p, err)
		return
	}
	log.Printf("room: wrote to %v %q\n", p, bs)
}

// nextRound starts the next round, using snippet r. nextRound is not
// concurrency safe.
func (r *room) nextRound(nextSnippet snippet.Snippet) {
	r.snippet = nextSnippet
	oldId := r.roundId
	r.roundId++
	r.playerTypedFirstCorrectChar = make(map[*player]time.Time)

	bs, err := json.Marshal(outgoingMsg{
		NewRoundMsg: &NewRoundOutgoingMsg{
			Snippet:    r.snippet.Code,
			NewRoundId: r.roundId,
			RoundId:    oldId,
		},
	})
	if err != nil {
		panic("marshalling a outgoingMsg should never result in an error")
	}
	for pp := range r.players {
		r.sendTo(pp, bs)
	}
}

func (r *room) handlePlayerTypedCorrectChars(p *player, correctChars int) {
	r.Lock()
	defer r.Unlock()

	if correctChars == 1 {
		r.playerTypedFirstCorrectChar[p] = time.Now()
	}

	if correctChars == len(r.snippet.Code) {
		seconds := time.Since(r.playerTypedFirstCorrectChar[p]).Seconds()
		characters := utf8.RuneCountInString(r.snippet.Code)
		chatMessageContent := fmt.Sprintf(
			"%s won the round, he or she typed it in %.2f seconds at %.2f characters per second!",
			p,
			seconds,
			float64(characters)/seconds,
		)
		bs, err := json.Marshal(outgoingMsg{
			ChatMessageMsg: &ChatMessageOutgoingMsg{
				Sender:  "liracer",
				Content: chatMessageContent,
			},
		})
		if err != nil {
			panic("marshalling a outgoingMsg should never result in an error")
		}
		for pp := range r.players {
			r.sendTo(pp, bs)
		}

		r.nextRound(snippet.Random())
		return
	}
	bs, err := json.Marshal(
		outgoingMsg{
			OpponentCorrectCharsMsg: &OpponentCorrectCharsOutgoingMsg{
				OpponentID:   p.id,
				CorrectChars: correctChars,
				RoundId:      r.roundId,
			},
		},
	)
	if err != nil {
		panic("marshalling a outgoingMsg should never result in an error")
	}
	for pp := range r.players {
		if pp != p {
			r.sendTo(pp, bs)
		}
	}
}

func (r *room) handlePlayerSentChatMessage(p *player, content string) {
	r.Lock()
	defer r.Unlock()

	if strings.HasPrefix(content, "/new-round-with-snippet") {
		snippetName := content[len("/new-round-with-snippet "):]
		s := snippet.Get(snippetName)
		if s == nil {
			content := fmt.Sprintf("there is no snippet with name %q", snippetName)
			bs, err := json.Marshal(outgoingMsg{
				ChatMessageMsg: &ChatMessageOutgoingMsg{
					Content: content,
					Sender:  "liracer",
				},
			})
			if err != nil {
				panic("marshalling a outgoingMsg should never result in an error")
			}
			r.sendTo(p, bs)
		} else {
			content := fmt.Sprintf("%v started a new round with snippet %s.", p, s.Name)
			bs, err := json.Marshal(outgoingMsg{
				ChatMessageMsg: &ChatMessageOutgoingMsg{
					Content: content,
					Sender:  "liracer",
				},
			})
			if err != nil {
				panic("marshalling a outgoingMsg should never result in an error")
			}
			for p := range r.players {
				r.sendTo(p, bs)
			}
			r.nextRound(*s)
		}
		return
	}

	bs, err := json.Marshal(outgoingMsg{
		ChatMessageMsg: &ChatMessageOutgoingMsg{
			Content: content,
			Sender:  p.String(),
		},
	})
	if err != nil {
		panic("marshalling a outgoingMsg should never result in an error")
	}

	for p := range r.players {
		r.sendTo(p, bs)
	}
}

func (r *room) handlePlayerJoined(p *player) {
	r.Lock()
	defer r.Unlock()

	r.players[p] = struct{}{}
	log.Printf("room: %v joined, there are now %d players\n", p, len(r.players))

	bs, err := json.Marshal(
		outgoingMsg{
			NewRoundMsg: &NewRoundOutgoingMsg{
				Snippet:    r.snippet.Code,
				NewRoundId: r.roundId,
				RoundId:    0,
			},
		},
	)
	if err != nil {
		panic("marshalling a outgoingMsg should never result in an error")
	}
	r.sendTo(p, bs)

	for pp := range r.players {
		bs, err := json.Marshal(
			outgoingMsg{
				ChatMessageMsg: &ChatMessageOutgoingMsg{
					Content: fmt.Sprintf("%s joined the room", p),
					Sender:  "liracer",
				},
			},
		)
		if err != nil {
			panic("marshalling a outgoingMsg should never result in an error")
		}
		r.sendTo(pp, bs)
	}
}

func (r *room) handlePlayerLeft(p *player) {
	r.Lock()
	defer r.Unlock()

	delete(r.players, p)
	log.Printf("room: %v left, there are now %d players\n", p, len(r.players))

	for pp := range r.players {
		bs, err := json.Marshal(
			outgoingMsg{
				ChatMessageMsg: &ChatMessageOutgoingMsg{
					Content: fmt.Sprintf("%s left the room", p),
					Sender:  "liracer",
				},
			},
		)
		if err != nil {
			panic("marshalling a outgoingMsg should never result in an error")
		}
		r.sendTo(pp, bs)
	}
}
