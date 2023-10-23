package main

import (
	"context"
	"fmt"

	pb "github.com/NicoPUSM/PruebaLab1/Prueba-Lab1/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	var estado string

	if err != nil {
		panic("no se puede conectar con el servidor" + err.Error())
	}
	for {
		fmt.Print("Escribe infectado o muerto: ")
		fmt.Scan(&estado)
		fmt.Println(estado)

		serviceClient := pb.NewMensajeServiceClient(conn)

		res, err := serviceClient.CreateLista(context.Background(), &pb.ConsultarLista{
			Estado: &pb.Estado{
				Nombre: estado,
			},
		})

		if err != nil {
			panic("no se creo el mensaje" + err.Error())
		}

		for _, nombre := range res.Estadoid {
			fmt.Println(nombre)
		}

		fmt.Println(" ")
	}
}
