package msg

// OutgoingMsg is used to marshal outgoing messages.
type OutgoingMsg struct {
	// If the SetRoomStateMsg field is non-nil, this field is used to tell a
	// player to set their snippet  to SetRoomStateMsg.Snippet.
	SetRoomStateMsg *SetRoomStateOutgoingMsg
}

// SetRoomStateOutgoingMsg is used exclusively as an optional field of
// OutgoingMsg. Therefore the documentation for it lives there.
type SetRoomStateOutgoingMsg struct {
	Snippet string
}
