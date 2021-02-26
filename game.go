// game type
package main

type gameId string

// game represents a game and provides concurrency safe methods of registering
// and unregistering players to itself. It also provides methods for interacting
// with the players in the game.
type game struct {
	players map[*player]struct{} // preferably we should not expose this state
	snippet string
}

func (g *game) register(p *player) { // NB: must be concurrency saye
}

func (g *game) unregister(p *player) { // NB: must be concurrency saye
}

// func (g *game) sendToAllPlayer(bs []byte) {
// 	// WARNING: this will in some cases send bs even to players which have
// 	//          already left the game. It the job of the client to
// 	//          invalidate these messages.
// 	for p := range g.players {
// 		p.send <- bs
// 	}
// }

// func (g *game) sendToAllPlayersExcept(p *player) {
// 	// WARNING: see WARNING in sendToAllPlayers
// 	for pp := range g.players {
// 		if p != pp {
// 			p.send <- bs
// 		}
// 	}
// }
