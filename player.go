// player type and newPlayer function.
package main

type playerId int

// player represents a player.
type player struct {
	// id identifies player.
	id playerId
	// gid identifies player's current game. If gid == "" player is not in a
	// game.
	gid gameId
	// send is a channel for writing to player's underlying connection.
	send chan []byte
}

// playerIdEmitter generates playerIds. It first produces playerId(1). Thereafter
// it always produces playerIds one greater than the previous, I.e playerId(1),
// playerId(2), ..., playerId(n).
var playerIdEmitter = make(chan playerId)

func init() {
	// No cleanup is necessary
	go func() {
		var nextId playerId = 1
		for {
			playerIdEmitter <- nextId
			nextId++
		}
	}()
}

// newPlayer creates a player with a unique id. It initializes send, but does not
// start reading from it. That job is reserved for the caller.
func newPlayer() *player {
	return &player{
		id:   <-playerIdEmitter,
		gid:  gameId(""),
		send: make(chan []byte),
	}
}
