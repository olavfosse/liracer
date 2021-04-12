package main

// incomingMsg is used to unmarshal incoming messages.
type incomingMsg struct {
	// If the CorrectCharsMsg field is non-nil, this field is used by the
	// players to broadcast that they have typed CorrectCharsMsg.CorrectChars
	// characters correctly to the other players in their room.
	// CorrectCharsMsg.RoundId is used for invalidating the message in case it
	// "belongs" to a previous round.
	CorrectCharsMsg *CorrectCharsIncomingMsg
}

// CorrectCharsIncomingMsg is used exclusively as an optional field of
// incomingMsg. Therefore the documentation for it lives there.
type CorrectCharsIncomingMsg struct {
	CorrectChars int
	RoundId      roundId
}
