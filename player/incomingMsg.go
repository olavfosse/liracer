package player

// incomingMsg is used to unmarshal incoming messages.
type incomingMsg struct {
	// If the JoinRoomMsg field is non-nil, this field is used by the players
	// to join the room.
	JoinRoomMsg *JoinRoomIncomingMsg
	// If the CorrectCharsMsg field is non-nil, this field is used by the
	// players to broadcast that they have typed CorrectCharsMsg.CorrectChars
	// characters correctly to the other players in their room.
	CorrectCharsMsg *CorrectCharsIncomingMsg
}

// JoinRoomIncomingMsg is used exclusively as an optional field of incomingMsg.
// Therefore the documentation for it lives there.
type JoinRoomIncomingMsg struct{}

// CorrectCharsIncomingMsg is used exclusively as an optional field of
// incomingMsg. Therefore the documentation for it lives there.
type CorrectCharsIncomingMsg struct {
	CorrectChars int
}
