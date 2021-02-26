// Types for marshalling/unmarshalling JSON encoded data.
package main

// incomingMsg is used to unmarshal incoming messages.
type incomingMsg struct {
	// If the JoinGameMsg field is non-nil, this field is used by the players
	// to join/create the game by id JoinGameMsg.GameId.
	JoinGameMsg *JoinGameIncomingMsg
	// If the CorrectCharsMsg field is non-nil, this field is used by the
	// players to inform the server that they have changed their number of
	// CorrectChars to CorrectCharsMsg.CorrectChars.
	CorrectCharsMsg *CorrectCharsIncomingMsg
}

// JoinGameIncomingMsg is used exclusively as an optional field of incomingMsg.
// Therefore the documentation for it lives there.
type JoinGameIncomingMsg struct {
	GameId gameId
}

// CorrectCharsIncomingMsg is used exclusively as an optional field of
// incomingMsg. Therefore the documentation for it lives there.
type CorrectCharsIncomingMsg struct {
	CorrectChars int
}

// outgoingMsg is used to marshal outgoing messages.
type outgoingMsg struct {
	// The GameId field describes which game the message is sent from. This
	// is used by the player to invalidate messages sent from other games than
	// the player's current game.
	GameId gameId
	// If the SetGameStateMsg field is non-nil, this field is used to tell a
	// player to set their roundId to SetGameStateMsg.RoundId and set their
	// snippet to SetGameStateMsg.Snippet.
	SetGameStateMsg *SetGameStateOutgoingMsg
	// // If the OpponentCorrectChars field is non-nil, this field is used to
	// // inform a player that his opponent, with id OpponentId, has changed
	// // their number of CorrectChars.
	// OpponentCorrectChars *struct {
	// 	OpponentId   playerId
	// 	CorrectChars int
	// }
}

// SetGameStateOutgoingMsg is used exclusively as an optional field of
// outgoingMsg. Therefore the documentation for it lives there.
type SetGameStateOutgoingMsg struct {
	RoundId roundId
	Snippet string
}
