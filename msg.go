// Types for marshalling/unmarshalling JSON data.
package main

type incomingMsg struct {
	JoinGameMsg     *JoinGameMsg
	CorrectCharsMsg *CorrectCharsMsg
}

type JoinGameMsg struct {
	GameId string
}

type CorrectCharsMsg struct {
	CorrectChars int
}

type outgoingMsg struct {
	SnippetMsg *SnippetMsg
}

type SnippetMsg struct {
	Snippet string
}
