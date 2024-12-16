package app

import (
	"log"

	"SepGateway/config"
)

func StartApplication() {
	// Cargar variables de entorno
	config.LoadEnv()

	// Obtener configuración del socket
	host := config.GetEnv("SOCKET_HOST", "localhost")
	port := config.GetEnv("SOCKET_PORT", "7776")

	// Crear cliente de socket
	client, err := config.NewSocketClient(host, port)
	if err != nil {
		log.Fatalf("Error al conectar con EscuelaGateway: %v", err)
	}
	defer client.Close()

	// Crear delegado de mensajes
	delegate := config.NewMessageDelegate(client)

	// Iniciar la recepción y manejo de mensajes
	delegate.ReceiveAndHandle()
}
