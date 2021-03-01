// Types for marshalling/unmarshalling JSON encoded data.
package player

// incomingMsg is used to unmarshal incoming messages.
type incomingMsg struct {
	// If the JoinGameMsg field is non-nil, this field is used by the players
	// to join the game.
	JoinGameMsg *JoinGameIncomingMsg
}

// JoinGameIncomingMsg is used exclusively as an optional field of incomingMsg.
// Therefore the documentation for it lives there.
type JoinGameIncomingMsg struct{}

// outgoingMsg is used to marshal outgoing messages.
type outgoingMsg struct {
	// If the SetGameStateMsg field is non-nil, this field is used to tell a
	// player to set their snippet  to SetGameStateMsg.Snippet.
	SetGameStateMsg *SetGameStateOutgoingMsg
}

// SetGameStateOutgoingMsg is used exclusively as an optional field of
// outgoingMsg. Therefore the documentation for it lives there.
type SetGameStateOutgoingMsg struct {
	Snippet string
}
