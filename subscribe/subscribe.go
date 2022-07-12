package sub

import (
	log "HermesMQ/logging"
	"HermesMQ/server"
)

var (
	listening ListenSubscribe = ListenSubscribe{
		Listening: []server.ActiveConnection{},
	}
)

type ListenSubscribe struct {
	Listening []server.ActiveConnection
}

func (ls *ListenSubscribe) Add(s server.ActiveConnection) {
	log.Infof("################Add novo ouvinte ", s.Name)
	ls.Listening = append(ls.Listening, s)
	log.Infof("Clientes connectados  [ %d ] ", len(listening.Listening))
}

type Subscribe struct {
	Name string
	Data string
}
