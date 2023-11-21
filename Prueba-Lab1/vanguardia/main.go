package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	pb "github.com/NicoPUSM/PruebaLab1/Prueba-Lab1/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("dist088:50051", grpc.WithInsecure())

	if err != nil {
		panic("no se puede conectar con el servidor" + err.Error())
	}
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Escribe el sector que consultar: ")

		scanner.Scan()
		estado := scanner.Text()

		if strings.ToLower(estado) == "exit" {
			break
		}

		serviceClient := pb.NewMensajeServiceClient(conn)

		res, err := serviceClient.CreateLista(context.Background(), &pb.ConsultarLista{
			Estado: &pb.Estado{
				Nombre: estado,
			},
		})

		if err != nil {
			panic("no se creo el mensaje" + err.Error())
		}

		for _, lista := range res.Estadoid {
			fmt.Println(lista)
		}

		fmt.Println(" ")
	}
}
