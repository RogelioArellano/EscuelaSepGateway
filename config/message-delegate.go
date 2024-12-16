package config

import (
	"encoding/json"
	"log"

	"SepGateway/models"
)

type MessageDelegate struct {
	client *SocketClient
}

// NewMessageDelegate crea una nueva instancia de MessageDelegate
func NewMessageDelegate(client *SocketClient) *MessageDelegate {
	return &MessageDelegate{client: client}
}

// ReceiveAndHandle recibe y maneja mensajes de manera continua
func (md *MessageDelegate) ReceiveAndHandle() {
	for {
		messageBytes, err := md.client.Receive()
		if err != nil {
			log.Printf("Error al recibir mensaje: %v\n", err)
			// Aquí podrías intentar reconectar o manejar el error según tus necesidades
			break
		}

		// Deserializar el mensaje recibido
		var mensaje models.Message // Define tu estructura de mensaje
		err = json.Unmarshal(messageBytes, &mensaje)
		if err != nil {
			log.Printf("Error al deserializar mensaje: %v\n", err)
			continue
		}

		// Procesar el mensaje
		md.handleMessage(mensaje)
	}
}

func (md *MessageDelegate) handleMessage(mensaje models.Message) {
	err := EnviarAMicroservicio(mensaje)
	if err != nil {
		log.Printf("Error al enviar mensaje al microservicio: %v\n", err)
	}
}
