package main

// outgoingMsg is used to marshal outgoing messages.
type outgoingMsg struct {
	// If the NewRoundMsg field is non-nil, this field is used to tell a player
	// that a new round has begun with snippet NewRoundMsg.Snippet and id
	// NewRoundMsg.NewRoundId. NewRoundMsg.RoundId is used for invalidating the
	// message in case it "belongs" to a previous round.
	NewRoundMsg *NewRoundOutgoingMsg
	// If the OpponentCorrectCharsMsg field is non-nil, this field is used
	// to a player that their opponent with ID
	// OpponentCorrectCharsMsg.OpponentID has typed
	// OpponentCorrectCharsMsg.CorrectChars characters correctly.
	// OpponentCorrectCharsMsg.RoundId is used for invalidating the message in
	// case it "belongs" to a previous round.
	OpponentCorrectCharsMsg *OpponentCorrectCharsIncomingMsg
}

// NewRoundOutgoingMsg is used exclusively as an optional field of outgoingMsg.
// Therefore the documentation for it lives there.
type NewRoundOutgoingMsg struct {
	Snippet    string
	NewRoundId roundId
	RoundId    roundId
}

// OpponentCorrectCharsIncomingMsg is used exclusively as an optional field of
// outgoingMsg. Therefore the documentation for it lives there.
type OpponentCorrectCharsIncomingMsg struct {
	OpponentID   id
	CorrectChars int
	RoundId      roundId
}
