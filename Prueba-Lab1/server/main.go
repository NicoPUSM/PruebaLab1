package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/NicoPUSM/PruebaLab1/Prueba-Lab1/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnsafeMensajeServiceServer
}

func (s *server) Create(ctx context.Context, req *pb.Crearmensaje) (*pb.Respuestamensaje, error) {
	fmt.Printf("creando mensaje " + req.Mensaje.Nombre)

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
