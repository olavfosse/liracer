// vi bruker roundId og gameId til å invalidere meldinger som kan være ugyldig.

// Types for marshalling/unmarshalling JSON data.
package main

type roundId int

// incomingMsg is the type we unmarshal incoming websocket messages into.
type incomingMsg struct {
	JoinGameIncomingMsg     *JoinGameIncomingMsg
	CorrectCharsIncomingMsg *CorrectCharsIncomingMsg
}

// JoinGameIncomingMsg is used to join a game.
type JoinGameIncomingMsg struct {
	GameId string
}

// CorrectCharsIncomingMsg is used to inform the server of how many characters
// the player has typed correctly
type CorrectCharsIncomingMsg struct {
	RoundId      roundId
	CorrectChars int
}

// outgoingMsg is the type we use for marshalling outgoing websocket messages.
type outgoingMsg struct {
	GameId               gameId
	GameStateOutgoingMsg *GameStateOutgoingMsg
}

// // ChatMessageOutgoingMsg is used to relay a ChatMessageIncomingMsg to the other
// // players in the game.
// type ChatMessageOutgoingMsg struct {
// 	Sender  playerId
// 	Content string
// }

// GameStateOutgoingMsg is used to bring players the game state after they
// join/create a game or after a new round has started.
type GameStateOutgoingMsg struct {
	RoundId roundId
	Snippet string
	// OpponentCorrectChars map[playerId]int
}

// CorrectCharsOutgoingMsg is used to inform the players that player with id
// Opponent has sent a CorrectCharsIncomingMsg
type CorrectCharsOutgoingMsg struct {
	RoundId      roundId
	Opponent     playerId
	CorrectChars int
}
