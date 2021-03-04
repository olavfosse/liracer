package room

var Singleton *Room

func init() {
	Singleton = New()
}
