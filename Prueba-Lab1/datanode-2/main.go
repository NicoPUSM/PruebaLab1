package main

import (
	"context"
	"fmt"
	"net"
	"os"

	pb "github.com/NicoPUSM/PruebaLab1/Prueba-Lab1/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnsafeDataServiceServer
}

func (s *server) Create(ctx context.Context, req *pb.Crearmensaje) (*pb.Respuestamensaje, error) {
	archivo, err := os.OpenFile("DATA.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println("Error al crear archivo", err)
	}
	defer archivo.Close()

	fmt.Println("Recibio a " + req.Mensaje.Nombre)

	_, err = archivo.WriteString(req.Mensaje.Nombre + "\n")

	if err != nil {
		fmt.Println("Error al escribir en el archivo", err)
	}

	return &pb.Respuestamensaje{
		Mensajeid: req.Mensaje.Nombre,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50053")
	if err != nil {
		panic("no se creo la conexión tcp " + err.Error())
	}

	serv := grpc.NewServer()
	pb.RegisterDataServiceServer(serv, &server{})

	if err = serv.Serve(listener); err != nil {
		panic("no se inicio el servidor " + err.Error())
	}
}
