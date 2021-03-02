package msg

// OutgoingMsg is used to marshal outgoing messages.
type OutgoingMsg struct {
	// If the SetGameStateMsg field is non-nil, this field is used to tell a
	// player to set their snippet  to SetGameStateMsg.Snippet.
	SetGameStateMsg *SetGameStateOutgoingMsg
}

// SetGameStateOutgoingMsg is used exclusively as an optional field of
// OutgoingMsg. Therefore the documentation for it lives there.
type SetGameStateOutgoingMsg struct {
	Snippet string
}
