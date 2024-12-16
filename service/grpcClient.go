package service

import (
	"context"
	"log"
	"time"

	"SepGateway/models"

	"google.golang.org/grpc"
)

func EnviarAMicroservicio(mensaje models.Message) error {
	conn, err := grpc.Dial("direccion_microservicio:puerto", grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := pb.NewSEPServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	respuesta, err := client.EnviarMensaje(ctx, &mensaje)
	if err != nil {
		return err
	}

	log.Printf("Respuesta del microservicio: %v", respuesta)
	return nil
}
