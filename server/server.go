package server

import (
	log "HermesMQ/logging"
	"bufio"
	"encoding/json"
	"math/rand"
	"net"
	"strings"
	"time"
)

func StartServer(server Server) {
	server.ListenAndAccept()
}

type Server struct {
	Port             string
	IpAddress        string
	Protocol         string
	Description      string
	HandleConnection func(conn ActiveConnection)
}

//iniciar o servidor TCP e criar a conexão que fica ouvindo os novos clientes
//e invocada um gorotine para tratar as requeisções dos clientes
func (server *Server) ListenAndAccept() {
	var connection net.Conn
	listener, err := net.Listen(server.Listen())
	if err != nil {
		panic("Deu ruim ao  iniciar o servidor !" + err.Error() + "| config " + server.Socket() + "\n")
	}
	rand.Seed(time.Now().Unix()) //nao entendi bem o efeito disso!

	log.Infof("**** Aceitando conexoes [%s] Socket [%s]", server.Description, server.Socket())
	defer listener.Close()
	for {
		connection, err = listener.Accept()
		log.Infof("Nova conexao aceita :  %s ", connection.RemoteAddr().String())
		if err != nil {
			panic("Erro ao aceitar conexoes  " + err.Error())
		}

		ac := makeActiveConnection(connection)
		ac.Conn = connection
		//go topic.HandleConnection(topic.Topic{Meta: topic.TopicMeta{Conn: connection}})
		server.HandleConnection(ac)

		log.Infof("Conexao aceita com sucesso %s ", connection.RemoteAddr().String())
	}

}

//Recebe o primeiro envio de dados para gravar as informaoes necessarias
//para identificar o topico, e metadados de configurações do topico/subscrição (que atualmente nao tem nenhum)
func makeActiveConnection(connection net.Conn) (activeConnection ActiveConnection) {
	netData, _, err := bufio.NewReader(connection).ReadLine()
	if err != nil {
		log.Error(err.Error())
		return
	}
	tempDataMessage := strings.TrimSpace(string(netData))
	err = json.Unmarshal([]byte(tempDataMessage), &activeConnection)
	if err != nil {
		log.Errorf("Erro unmarshal activeConnection %s", err.Error())
	}

	log.Infof("Conexao ativa para objeto %s", activeConnection.Name)
	return
}

type ActiveConnection struct {
	Name string
	Conn net.Conn
}

type InstanceServer struct {
	Server     *Server
	Listener   net.Listener
	Connection net.Conn
}

func (s Server) Socket() string {
	return s.IpAddress + ":" + s.Port
}

func (s Server) Listen() (string, string) {
	return s.Protocol, s.IpAddress + ":" + s.Port
}
