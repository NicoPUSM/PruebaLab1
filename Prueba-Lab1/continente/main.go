package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	pb "github.com/NicoPUSM/PruebaLab1/Prueba-Lab1/proto"
	"google.golang.org/grpc"
)

func main() {
	content, err := os.ReadFile("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		panic("no se puede conectar con el servidor" + err.Error())
	}

	serviceClient := pb.NewMensajeServiceClient(conn)

	res, err := serviceClient.Create(context.Background(), &pb.Crearmensaje{
		Mensaje: &pb.Mensaje{
			Nombre: "Pedro",
		},
	})

	if err != nil {
		panic("no se creo el mensaje" + err.Error())
	}

	fmt.Println(res.Mensajeid)
}
