package game

var Singleton *Game

func init() {
	Singleton = New()
}
