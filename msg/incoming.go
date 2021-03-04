package msg

// IncomingMsg is used to unmarshal incoming messages.
type IncomingMsg struct {
	// If the JoinRoomMsg field is non-nil, this field is used by the players
	// to join the room.
	JoinRoomMsg *JoinRoomIncomingMsg
}

// JoinRoomIncomingMsg is used exclusively as an optional field of IncomingMsg.
// Therefore the documentation for it lives there.
type JoinRoomIncomingMsg struct{}
