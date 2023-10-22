package main

import (
	"context"
	"fmt"

	"github.com/NicoPUSM/PruebaLab1/Prueba-Lab1/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure()) // Conectar al puerto 50052
	var estado string

	if err != nil {
		panic("no se puede conectar con el servidor" + err.Error())
	}

	fmt.Print("Escribe afectado o muerto: ")
	fmt.Scan(&estado)
	fmt.Println(estado)

	serviceClient := proto.NewMensajeServiceClient(conn)

	// Llama a la función correcta en el servicio, que en este caso parece ser "Create".
	res, err := serviceClient.Create(context.Background(), &proto.Crearmensaje{
		Mensaje: &proto.Mensaje{
			Nombre: estado,
		},
	})

	if err != nil {
		panic("no se creó el mensaje" + err.Error())
	}

	fmt.Println(res.Mensajeid)
}
