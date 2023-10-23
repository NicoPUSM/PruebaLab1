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

func (s *server) CreateLista(ctx context.Context, req *pb.ConsultarLista) (*pb.RespuestaLista, error) {
	return nil, fmt.Errorf("CreateLista is not implemented")
}

func main() {
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		panic("no se creo la conexión tcp " + err.Error())
	}

	serv := grpc.NewServer()
	pb.RegisterMensajeServiceServer(serv, &server{})

	if err = serv.Serve(listener); err != nil {
		panic("no se inicio el servidor " + err.Error())
	}
}
