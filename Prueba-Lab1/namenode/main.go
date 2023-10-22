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
	fmt.Println("Recibió a " + req.Mensaje.Nombre)

	archivo, err := os.OpenFile("DATA.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error al crear archivo", err)
	}
	defer archivo.Close()

	palabras := strings.Split(req.Mensaje.Nombre, " ")

	var datanode string

	cadena := strconv.Itoa(contador)

	if string(palabras[1][0]) <= "M" {
		datanode = "1"
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
		datanode = "2"
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

	_, err = archivo.WriteString(cadena + " " + datanode + " " + palabras[2] + "\n")

	if err != nil {
		fmt.Println("Error al escribir en el archivo", err)
	}

	return &pb.Respuestamensaje{
		Mensajeid: req.Mensaje.Nombre,
	}, nil
}

func (s *server) ConsultarEstado(ctx context.Context, req *pb.ConsultarEstadoRequest) (*pb.ConsultarEstadoResponse, error) {
	// Lógica para consultar el estado aquí.
	// Por ahora, simplemente retornamos una lista de ejemplo.
	resultados := []string{"Estado1", "Estado2", "Estado3"}

	return &pb.ConsultarEstadoResponse{
		Resultados: resultados,
	}, nil
}

func main() {
	contador = 0

	// Iniciar el servidor para el servicio de mensajes en el puerto 50051.
	listenerMsg, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic("No se pudo crear la conexión tcp " + err.Error())
	}

	servMsg := grpc.NewServer()
	pb.RegisterMensajeServiceServer(servMsg, &server{})

	if err = servMsg.Serve(listenerMsg); err != nil {
		panic("No se pudo iniciar el servidor de mensajes " + err.Error())
	}

	listenerMsg.Close()

	listenerState, err := net.Listen("tcp", ":50054")
	if err != nil {
		panic("No se pudo crear la conexión tcp para el servicio de estado " + err.Error())
	}

	servState := grpc.NewServer()
	pb.RegisterMensajeServiceServer(servState, &server{})

	if err = servState.Serve(listenerState); err != nil {
		panic("No se pudo iniciar el servidor de estado " + err.Error())
	}

	listenerState.Close()
}
