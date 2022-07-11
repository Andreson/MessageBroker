package server

import (
	log "HermesMQ/logging"
	topic "HermesMQ/topic"
	"math/rand"
	"net"
	"time"
)

func StartServer(server Server) {
	server.ListenAndAccept()
}

type Server struct {
	Port      string
	IpAddress string
	Protocol  string
}

type InstanceServer struct {
	Server     *Server
	Listener   net.Listener
	Connection net.Conn
}

func (s Server) Socket() string {
	return s.IpAddress + ":" + s.Port
}

func (s Server) ConfigListen() (string, string) {
	return s.Protocol, s.IpAddress + ":" + s.Port
}

//iniciar o servidor TCP e criar a conexão que fica ouvindo os novos clientes
//e invocada um gorotine para tratar as requeisções dos clientes
func (server Server) ListenAndAccept() {
	var connection net.Conn
	listener, err := net.Listen(server.ConfigListen())
	if err != nil {
		panic("Deu ruim ap iniciar o servidor !" + err.Error() + "| config " + server.Socket() + "\n")
	}
	rand.Seed(time.Now().Unix()) //nao entendi bem o efeito disso!

	log.Info("Aceitando conexoes ****")
	defer listener.Close()
	for {
		connection, err = listener.Accept()
		log.Infof("Nova conexao aceita :  %s ", connection.RemoteAddr().String())
		if err != nil {
			panic("Erro ao aceitar conexoes  " + err.Error())
		}

		go topic.HandleConnection(topic.Topic{Meta: topic.TopicMeta{Conn: connection}})

		log.Infof("Conexao aceita com sucesso %s ", connection.RemoteAddr().String())
	}

}
