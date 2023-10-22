package main

import (
	"context"
	"fmt"

	pb "github.com/NicoPUSM/PruebaLab1/Prueba-Lab1/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50054", grpc.WithInsecure())
	var estado string

	if err != nil {
		panic("no se puede conectar con el servidor" + err.Error())
	}

	fmt.Print("Escribe afectado o muerto: ")
	fmt.Scan(&estado)
	fmt.Println(estado)

	serviceClient := pb.NewMensajeServiceClient(conn)

	res, err := serviceClient.Create(context.Background(), &pb.Crearmensaje{
		Mensaje: &pb.Mensaje{
			Nombre: estado,
		},
	})

	if err != nil {
		panic("no se creo el mensaje" + err.Error())
	}

	fmt.Println(res.Mensajeid)

}
