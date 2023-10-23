package main

import (
	"bufio"
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	pb "github.com/NicoPUSM/PruebaLab1/Prueba-Lab1/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		panic("no se puede conectar con el servidor" + err.Error())
	}

	content, err := os.Open("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(content)

	for scanner.Scan() {
		line := scanner.Text()

		rand.Seed(time.Now().UnixNano())
		randomValue := rand.Float64()
		var resultado string

		if randomValue < 0.55 {
			resultado = "infectada"
		} else {
			resultado = "muerta"
		}

		line = line + " " + resultado

		fmt.Println(line)

		serviceClient := pb.NewMensajeServiceClient(conn)

		res, err := serviceClient.Create(context.Background(), &pb.Crearmensaje{
			Mensaje: &pb.Mensaje{
				Nombre: line,
			},
		})

		if err != nil {
			panic("no se creo el mensaje" + err.Error())
		}

		fmt.Println(res.Mensajeid)

	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	content.Close()

}
