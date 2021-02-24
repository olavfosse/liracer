package game

import (
	"github.com/fossegrim/play.liracer.org/player"
	"github.com/gorilla/websocket"
)

type Game struct {
	players []*websocket.Conn
	snippet string
}

func (g *Game) Register(p player.Player) {

}

func (g *Game) Unregister(p player.Player) {
}
