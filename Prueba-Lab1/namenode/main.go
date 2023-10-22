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

func main() {
	contador := 0

	go startDataNodeServer(":50051", &contador)
	go startNameNodeServer(":50054")

	select {}
}

func startDataNodeServer(address string, contador *int) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic("No se pudo crear la conexión tcp " + err.Error())
	}

	serv := grpc.NewServer()
	pb.RegisterMensajeServiceServer(serv, &server{})

	if err = serv.Serve(listener); err != nil {
		panic("No se pudo iniciar el servidor en " + address + ": " + err.Error())
	}
}

type nameNodeServer struct {
	pb.UnsafeMensajeServiceServer
}

func (s *nameNodeServer) Create(ctx context.Context, req *pb.Crearmensaje) (*pb.Respuestamensaje, error) {
	conn, err := grpc.Dial("localhost:50055", grpc.WithInsecure())
	if err != nil {
		fmt.Println("No se puede conectar con el DataNode: ", err)
		return nil, err
	}
	defer conn.Close()

	datanodeClient := pb.NewMensajeServiceClient(conn)

	res, err := datanodeClient.Create(ctx, &pb.Crearmensaje{
		Mensaje: &pb.Mensaje{
			Nombre: req.Mensaje.Nombre,
		},
	})
	if err != nil {
		panic("no se creo el mensaje" + err.Error())
	}

	fmt.Println(res.Mensajeid)

	return &pb.Respuestamensaje{
		Mensajeid: res.Mensajeid,
	}, nil
}

func startNameNodeServer(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic("No se pudo crear la conexión tcp " + err.Error())
	}

	serv := grpc.NewServer()
	pb.RegisterMensajeServiceServer(serv, &nameNodeServer{})

	if err = serv.Serve(listener); err != nil {
		panic("No se pudo iniciar el servidor en " + address + ": " + err.Error())
	}
}
