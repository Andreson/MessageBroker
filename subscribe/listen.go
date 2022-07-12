package sub

import (
	log "HermesMQ/logging"
	"HermesMQ/server"

	"encoding/json"
)

func HandleConnection(conn server.ActiveConnection) {

	log.Infof("Novo cliente connectado [ %s ] ", conn.Conn.RemoteAddr())
	listening.Add(conn)

}

func SendMessage(subscribe Subscribe) {
	var topicNotFound bool = true
	log.Infof("Tentando entregar mensagem para topico [ %s ] ", subscribe.Name)
	log.Infof("Clientes connectados  [ %d ] ", len(listening.Listening))
	for _, item := range listening.Listening {
		if item.Name == subscribe.Name {
			topicNotFound = false
			log.Infof("Topico %s encontrado para entrega de mensagem ", subscribe.Name)
			deliveryMenssage(item, subscribe)
			//break para permitir pattern pubsub, remover break
		}
	}
	if topicNotFound {
		//criar mecanismo de retry
	}
}

func deliveryMenssage(subConn server.ActiveConnection, data Subscribe) {
	messageJsonData, err := json.Marshal(data)

	if err != nil {
		log.Errorf("Ocorreu um erro ao serializar mensagem antes do envio ", err.Error())
	}
	subConn.Conn.Write(messageJsonData)
	log.Infof("Mensage menviada com sucesso")
}
