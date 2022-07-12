package sub

import (
	log "HermesMQ/logging"
	"HermesMQ/server"
	"HermesMQ/topic"
	"encoding/json"
)

func HandleConnection(conn server.ActiveConnection) {

	listening.Add(conn)

}

func SendMessage(topic topic.Topic) {
	var topicNotFound bool = true
	for _, subscribe := range listening.Listening {
		if subscribe.Name == topic.Name {
			topicNotFound = false
			deliveryMenssage(subscribe, topic.BuildMessage())
			//break para permitir pattern pubsub, remover break
		}
	}
	if topicNotFound {
		//criar mecanismo de retry
	}
}

func deliveryMenssage(subscribe server.ActiveConnection, data topic.TopicData) {
	messageJsonData, err := json.Marshal(data)

	if err != nil {
		log.Errorf("Ocorreu um erro ao serializar mensagem antes do envio ", err.Error())
	}
	subscribe.Conn.Write([]byte(messageJsonData))
	log.Debug("Mensage menviada com sucesso")
}
