package room

import (
	"fmt"
	"log"
	"time"

	"play.liracer.org/snippet"
)

// Room must never block.
// Room must do only cheap/fast work.
// IO is not Room's responsibility.

// A message to Room.
// JoinMessage | LeaveMessage
type RoomMessage interface{}

// A message to Player.
type PlayerMessage interface{}

type PlayerID int

type Join_RoomMessage struct {
	PlayerID           PlayerID
	PlayerMessageQueue chan<- PlayerMessage
}
type Leave_RoomMessage PlayerID
type ChatMessage_PlayerMessage struct {
	Sender  string
	Content string
}

type TypedCorrectChars_RoomMessage struct {
	PlayerID PlayerID
	RoundID  RoundID
	Chars    int
}

type TypedCorrectChars_PlayerMessage struct {
	PlayerID PlayerID
	Chars    int
}

type NewRound_PlayerMessage struct {
	Snippet    string
	NewRoundID RoundID
}

type ChatMessage_RoomMessage struct {
	PlayerId PlayerID
	Content  string
}
type RoundID int

func start(messageQueue <-chan RoomMessage, snippetSet *snippet.SnippetSet) {
	type player struct {
		messageQueue chan<- PlayerMessage
		correctChars int
		// if zero-value, player has not typed any correct chars
		typedFirstCorrectChar time.Time
	}

	players := map[PlayerID]*player{}
	currentSnippet := snippetSet.Random()
	roundID := RoundID(0)

	for {
		Message := <-messageQueue
		switch m := Message.(type) {
		case Join_RoomMessage:
			players[m.PlayerID] = &player{
				messageQueue: m.PlayerMessageQueue,
				correctChars: 0,
			}
			m.PlayerMessageQueue <- NewRound_PlayerMessage{
				Snippet:    currentSnippet.Code,
				NewRoundID: roundID,
			}
			for _, p := range players {
				// if id == m.id {
				// 	continue
				// }
				p.messageQueue <- ChatMessage_PlayerMessage{
					Sender:  "liracer",
					Content: fmt.Sprintf("Player %d joined the room", m.PlayerID),
				}
			}
		case Leave_RoomMessage:
			_, ok := players[PlayerID(m)]
			if !ok {
				log.Println("BUG: received Leave_RoomMessage from player who is not in the room")
			}
			delete(players, PlayerID(m))
			for _, p := range players {
				p.messageQueue <- ChatMessage_PlayerMessage{
					Sender:  "liracer",
					Content: fmt.Sprintf("Player %d left the room", PlayerID(m)),
				}
			}
		case TypedCorrectChars_RoomMessage:
			fmt.Println(m.PlayerID)
			p, ok := players[m.PlayerID]
			if !ok {
				log.Println("BUG: received TypedCorrectChars_RoomMessage from player who is not in the room")
			}
			// discard messages sent from last round
			if m.RoundID != roundID {
				continue
			}
			if p.typedFirstCorrectChar.IsZero() {
				println("zero")
				p.typedFirstCorrectChar = time.Now()
			}
			characters := len(currentSnippet.Code)
			if m.Chars != characters {
				for _, p := range players {
					p.messageQueue <- TypedCorrectChars_PlayerMessage{
						PlayerID: m.PlayerID,
						Chars:    m.Chars,
					}
				}
				continue
			}

			seconds := time.Since(p.typedFirstCorrectChar).Seconds()
			for _, p := range players {
				p.messageQueue <- ChatMessage_PlayerMessage{
					Sender: "liracer",
					Content: fmt.Sprintf(
						"Player %d won the round, he or she typed it in %.2f seconds at %.2f characters per second!",
						m.PlayerID,
						seconds,
						float64(characters)/seconds,
					),
				}
			}
			currentSnippet = snippetSet.Random()
			roundID++
			for _, p := range players {
				p.messageQueue <- NewRound_PlayerMessage{
					Snippet:    currentSnippet.Code,
					NewRoundID: roundID,
				}
			}
		case ChatMessage_RoomMessage:
			for _, p := range players {
				p.messageQueue <- ChatMessage_PlayerMessage{
					Sender:  fmt.Sprintf("Player %d", m.PlayerId),
					Content: m.Content,
				}
			}
		default:
			log.Println("BUG: unhandled Message:", m)
		}
	}
}

func Start() (messageQueue chan<- RoomMessage, err error) {
	snippetSet, err := snippet.ParseSnippetSet()
	if err != nil {
		return nil, err
	}
	messageQueue_ := make(chan RoomMessage, 1000)
	go start(messageQueue_, snippetSet)
	return messageQueue_, nil
}
