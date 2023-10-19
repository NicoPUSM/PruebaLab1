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
	pb.UnsafeMensajeServiceServer
}

func (s *server) Create(ctx context.Context, req *pb.Crearmensaje) (*pb.Respuestamensaje, error) {
	fmt.Println("Recibio a " + req.Mensaje.Nombre)

	archivo, err := os.Create("DATA.txt")
	if err != nil {
		fmt.Println("Error al crear archivo", err)
	}
	defer archivo.Close()

	_, err = archivo.WriteString(req.Mensaje.Nombre)

	if err != nil {
		fmt.Println("Erroe al escribir en el archivo", err)
	}

	return &pb.Respuestamensaje{
		Mensajeid: req.Mensaje.Nombre,
	}, nil
}

func main() {
	listner, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic("no se creo la conexion tcp " + err.Error())
	}

	serv := grpc.NewServer()

	pb.RegisterMensajeServiceServer(serv, &server{})

	if err = serv.Serve(listner); err != nil {
		panic("no se inicio el server " + err.Error())
	}
}
