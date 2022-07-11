package main

import (
	log "HermesMQ/logging"
	"HermesMQ/server"
	"fmt"
)

func main() {
	log.Info("######### Hermes mensage Queue #########")
	server.StartServer(server.Server{Protocol: "tcp", Port: "9010"})

	fmt.Println("Servidor iniciado ")

}
