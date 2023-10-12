package main

import (
	"context"
	"fmt"

	pb "github.com/Sistemas-Distribuidos-2023-02/Grupo22-Laboratorio-1/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		panic("no se puede conectar con el servidor" + err.Error())
	}

	serviceClient := pb.NewMensajeServiceClient(conn)

	res, err := serviceClient.Create(context.Background(), &pb.Crearmensaje{
		mensaje: &pb.mensaje{
			region: generarID,
		},
	})

	if err != nil {
		panic("no se creo el mensaje" + err.Error())
	}
	
	fmt.Println(res.mensaje)
}
