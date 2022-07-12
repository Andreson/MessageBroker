package main

import (
	log "HermesMQ/logging"
	"HermesMQ/server"
	"HermesMQ/topic"
	"fmt"
)

func main() {
	log.Info("######### Hermes mensage Queue #########")

	server.StartServer(server.Server{Protocol: "tcp", Port: "9010", HandleConnection: topic.HandleConnection, Description: "Listener Topic connection"})

	//server.StartServer(server.Server{Protocol: "tcp", Port: "9020", HandleConnection: sub.HandleConnection, Description: "Listener Subscribe connection"})

	fmt.Println("Servidor iniciado ")

}
