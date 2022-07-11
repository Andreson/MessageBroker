package topic

import (
	log "HermesMQ/logging"
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"
)

type Topic struct {
	Meta TopicMeta
	Name string
	Data string
}

//strutura mensagem para ser persistida
type TopicData struct {
	TopicName    string
	Content      string
	ReceivedTime time.Time
	ReadTime     time.Time
	Status       string
	ClientSocket string
}

//metadados do topic para manipulação da conexao
type TopicMeta struct {
	Conn            net.Conn
	ClientSocket    string
	Name            string
	startConnection time.Time
}

func (t Topic) BuildMessage() TopicData {

	return TopicData{TopicName: t.Name,
		Content:      t.Data,
		ReceivedTime: time.Now(),
		Status:       "Received",
		ClientSocket: t.Meta.ClientSocket}
}

//trata os dados recebidos dos clientes
func HandleConnection(t Topic) {
	log.Infof("Hanndler :  %s\n", t.Meta.Conn.RemoteAddr().String())
	for {
		netData, _, err := bufio.NewReader(t.Meta.Conn).ReadLine()
		if err != nil {
			fmt.Println(err)
			return
		}

		tempDataMessage := strings.TrimSpace(string(netData))

		if tempDataMessage == "STOP" {
			break
		} else {
			err := json.Unmarshal([]byte(tempDataMessage), &t)
			if err != nil {
				log.Infof("Erro unmarshal", err.Error())
			}
			err = WriteMessage(t)
			if err != nil {
				log.Errorf("Erro ao gravar a mensagem do topico[%s] [%s]\n", t.Name, err.Error())
			}
			log.Infof("Mensagem recebida: %s ", tempDataMessage)
		}

		t.Meta.Conn.Write([]byte("Conexao remota aberta em :" + t.Meta.Conn.RemoteAddr().String() + "\n"))
	}
	t.Meta.Conn.Close()
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}
