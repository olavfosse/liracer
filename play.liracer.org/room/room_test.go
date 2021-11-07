package room_test

import (
	"strings"
	"testing"

	"play.liracer.org/room"
)

func TestRoom(t *testing.T) {
	roomMessageQueue, err := room.Start()
	if err != nil {
		panic(err)
	}

	playerID := 2
	playerMessageQueue := make(chan room.PlayerMessage, 100)
	roomMessageQueue <- room.Join_RoomMessage{
		PlayerID:           room.PlayerID(playerID),
		PlayerMessageQueue: playerMessageQueue,
	}
	m := <-playerMessageQueue

	cm := m.(room.ChatMessage_PlayerMessage)
	if cm.Sender != "liracer" {
		t.Errorf("cm.Sender != \"liracer\"")
	}

	if !strings.Contains(cm.Content, "joined") {
		t.Errorf("%s does not contain \"joined\"", cm.Content)
	}
}
