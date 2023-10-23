package main

import (
	"bufio"
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
	fmt.Println("Solicitud de " + req.Mensaje.Nombre + " recibida, mensaje enviado: " + req.Mensaje.Nombre)

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

func (s *server) CreateLista(ctx context.Context, req *pb.ConsultarLista) (*pb.RespuestaLista, error) {
	lista := []string{}
	lista_1 := []string{}
	lista_2 := []string{}

	archivo, err := os.Open("DATA.txt")

	if err != nil {
		fmt.Println("Error al abrir archivo", err)
	}

	scanner := bufio.NewScanner(archivo)

	for scanner.Scan() {
		line := scanner.Text()
		div := strings.Fields(line)

		if div[2] == req.Estado.Nombre {
			if div[1] == "1" {
				lista_1 = append(lista_1, div[0])
			} else {
				lista_2 = append(lista_2, div[0])
			}
		}
	}

	archivo.Close()

	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())

	serviceClient := pb.NewMensajeServiceClient(conn)

	res, err := serviceClient.CreateMutuo(context.Background(), &pb.EnviarLista{
		Listado: &pb.Listado{
			Nombre: lista_1,
		},
	})

	if err != nil {
		panic("no se creo el mensaje" + err.Error())
	}

	lista = append(lista, res.Listaid...)

	conn1, err := grpc.Dial("localhost:50053", grpc.WithInsecure())

	serviceClient1 := pb.NewMensajeServiceClient(conn1)

	res1, err := serviceClient1.CreateMutuo(context.Background(), &pb.EnviarLista{
		Listado: &pb.Listado{
			Nombre: lista_2,
		},
	})

	if err != nil {
		panic("no se creo el mensaje" + err.Error())
	}

	lista = append(lista, res1.Listaid...)

	fmt.Println("Solicitud de " + req.Estado.Nombre + " recibida, mensaje enviado: ")
	fmt.Println(lista)

	return &pb.RespuestaLista{
		Estadoid: lista,
	}, nil
}

func (s *server) CreateMutuo(ctx context.Context, req *pb.EnviarLista) (*pb.RecibirLista, error) {
	return nil, fmt.Errorf("CreateMutuo is not implemented")
}

func main() {
	contador = 0

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
