package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	pb "github.com/NicoPUSM/PruebaLab1/Prueba-Lab1/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnsafeMensajeServiceServer
}

var contador int

func (s *server) Create(ctx context.Context, req *pb.Crearmensaje) (*pb.Respuestamensaje, error) {

	contador++
	fmt.Println("Recibio a " + req.Mensaje.Nombre)

	archivo, err := os.OpenFile("DATA.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error al crear archivo", err)
	}
	defer archivo.Close()

	palabras := strings.Split(req.Mensaje.Nombre, " ")

	var datanote string

	cadena := strconv.Itoa(contador)

	if string(palabras[1][0]) <= "M" {
		datanote = "1"
		conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
		if err != nil {
			fmt.Println("No se puede conectar con el DataNode: ", err)
			return nil, err
		}
		defer conn.Close()

		datanodeClient := pb.NewMensajeServiceClient(conn)

		_, err = datanodeClient.Create(ctx, &pb.Crearmensaje{
			Mensaje: &pb.Mensaje{
				Nombre: cadena + " " + palabras[0] + " " + palabras[1],
			},
		})
	} else {
		datanote = "2"
		conn, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
		if err != nil {
			fmt.Println("No se puede conectar con el DataNode: ", err)
			return nil, err
		}
		defer conn.Close()

		datanodeClient := pb.NewMensajeServiceClient(conn)

		_, err = datanodeClient.Create(ctx, &pb.Crearmensaje{
			Mensaje: &pb.Mensaje{
				Nombre: cadena + " " + palabras[0] + " " + palabras[1],
			},
		})
	}

	_, err = archivo.WriteString(cadena + " " + datanote + " " + palabras[2] + "\n")

	if err != nil {
		fmt.Println("Error al escribir en el archivo", err)
	}

	return &pb.Respuestamensaje{
		Mensajeid: req.Mensaje.Nombre,
	}, nil
}

func (s *server) ConsultarEstado(ctx context.Context, req *pb.ConsultarEstadoRequest) (*pb.ConsultarEstadoResponse, error) {

	resultados := []string{"hola"}

	return &pb.ConsultarEstadoResponse{
		Resultados: resultados,
	}, nil
}

func main() {
	contador = 0

	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		panic("No se pudo crear la conexión tcp " + err.Error())
	}

	serv := grpc.NewServer()
	pb.RegisterMensajeServiceServer(serv, &server{})

	if err = serv.Serve(listener); err != nil {
		panic("No se pudo iniciar el servidor " + err.Error())
	}

	listener.Close()
}
