package sub

import "HermesMQ/server"

var (
	listening ListenSubscribe = ListenSubscribe{
		Listening: []server.ActiveConnection{},
	}
)

type ListenSubscribe struct {
	Listening []server.ActiveConnection
}

func (ls ListenSubscribe) Add(s server.ActiveConnection) {
	ls.Listening = append(ls.Listening, s)
}
