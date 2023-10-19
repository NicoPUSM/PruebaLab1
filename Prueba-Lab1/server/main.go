package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/Grupo22-Laboratorio-2/prueba/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnsafeMensajeServiceServer
}

func (s *server) Create(ctx context.Context, req *pb.Crearmensaje) (*pb.Respuestamensaje, error) {
	fmt.Printf("creando mensaje " + req.Mensaje.Region)

	return &pb.Respuestamensaje{
		Mensajeid: req.Mensaje.Region,
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
