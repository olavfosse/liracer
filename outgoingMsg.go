package main

// outgoingMsg is used to marshal outgoing messages.
type outgoingMsg struct {
	// If the NewRoundMsg field is non-nil, this field is used to tell a player
	// that a new round has begun with snippet NewRoundMsg.Snippet.
	NewRoundMsg *NewRoundOutgoingMsg
	// If the OpponentCorrectCharsMsg field is non-nil, this field is used
	// to a player that their opponent with ID
	// OpponentCorrectCharsMsg.OpponentID has typed
	// OpponentCorrectCharsMsg.CorrectChars characters correctly.
	OpponentCorrectCharsMsg *OpponentCorrectCharsIncomingMsg
}

// NewRoundOutgoingMsg is used exclusively as an optional field of outgoingMsg.
// Therefore the documentation for it lives there.
type NewRoundOutgoingMsg struct {
	Snippet string
}

// OpponentCorrectCharsIncomingMsg is used exclusively as an optional field of
// outgoingMsg. Therefore the documentation for it lives there.
type OpponentCorrectCharsIncomingMsg struct {
	OpponentID   id
	CorrectChars int
}
