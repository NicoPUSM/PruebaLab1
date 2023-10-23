package main

import (
	"context"
	"fmt"

	"github.com/NicoPUSM/PruebaLab1/Prueba-Lab1/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50055", grpc.WithInsecure())
	var estado string

	if err != nil {
		panic("no se puede conectar con el servidor" + err.Error())
	}

	fmt.Print("Escribe afectado o muerto: ")
	fmt.Scan(&estado)
	fmt.Println(estado)

	serviceClient := proto.NewMensajeServiceClient(conn)

	res, err := serviceClient.ConsultarEstado(context.Background(), &proto.ConsultarEstadoRequest{
		Estado: &proto.Estado{
			Nombre: estado,
		},
	})

	if err != nil {
		panic("no se creó el mensaje" + err.Error())
	}

	fmt.Println(res.Resultados)
}
