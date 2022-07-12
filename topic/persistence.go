package topic

import (
	log "HermesMQ/logging"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
)

const (
	PATH_PERSIST_TOPICS = "/hermes/topic/data/"
)

//escreve mensagem no disco para ser persistida e recuperada em caso falha no servidor
func WriteMessage(topic Topic) error {
	//var file []byte
	topicPath := PATH_PERSIST_TOPICS + topic.Name + "/"
	if _, err := os.Stat(topicPath); err != nil {
		log.Infof("Criando diretorio do topico %s ", topicPath)
		err = os.Mkdir(topicPath, os.ModePerm)

		if err != nil {
			log.Errorf("Erro ao criar pasta para armazenamento do topico %s", topic.Name)
			return err
		}
	}
	file, err := json.MarshalIndent(topic.BuildMessage(), "", " ")
	if err != nil {
		log.Errorf("Erro ao deserializar mensagem  %s", err.Error())
		return err
	}
	fileName := topicPath + uuid.New().String() + ".json"
	err = ioutil.WriteFile(fileName, file, 0644)
	return err

}
