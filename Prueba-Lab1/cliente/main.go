package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	pb "github.com/NicoPUSM/PruebaLab1/Prueba-Lab1/proto"
	"google.golang.org/grpc"
)

func generarID() string {
	rand.Seed(time.Now().Unix())
	return "ID: " + strconv.Itoa(rand.Int())
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		panic("no se puede conectar con el servidor" + err.Error())
	}

	serviceClient := pb.NewMensajeServiceClient(conn)

	res, err := serviceClient.Create(context.Background(), &pb.Crearmensaje{
		Mensaje: &pb.Mensaje{
			Nombre: generarID(),
		},
	})

	if err != nil {
		panic("no se creo el mensaje" + err.Error())
	}

	fmt.Println(res.Mensajeid)
}
