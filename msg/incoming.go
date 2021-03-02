package msg

// IncomingMsg is used to unmarshal incoming messages.
type IncomingMsg struct {
	// If the JoinGameMsg field is non-nil, this field is used by the players
	// to join the game.
	JoinGameMsg *JoinGameIncomingMsg
}

// JoinGameIncomingMsg is used exclusively as an optional field of IncomingMsg.
// Therefore the documentation for it lives there.
type JoinGameIncomingMsg struct{}
