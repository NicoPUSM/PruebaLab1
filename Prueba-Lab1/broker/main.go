package main

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"time"

	pb "github.com/NicoPUSM/PruebaLab1/Prueba-Lab1/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnsafeMensajeServiceServer
}

func (s *server) Create(ctx context.Context, req *pb.Crearmensaje) (*pb.Respuestamensaje, error) {
	var direccionAleatoria string

	fmt.Println("Solicitud de " + req.Mensaje.Nombre + " recibida, mensaje enviado: " + req.Mensaje.Nombre)
	rand.Seed(time.Now().UnixNano())
	numeroAleatorio := rand.Intn(3)

	if numeroAleatorio == 0 {
		direccionAleatoria = "50052"

	} else if numeroAleatorio == 1 {
		direccionAleatoria = "50053"

	} else if numeroAleatorio == 2 {
		direccionAleatoria = "50054"

	}

	fmt.Println("Solicitud de " + req.Mensaje.Nombre + " recibida, mensaje enviado: " + req.Mensaje.Nombre)

	conn, err := grpc.Dial("localhost:"+direccionAleatoria, grpc.WithInsecure())

	if err != nil {
		panic("no se puede conectar con el servidor" + err.Error())
	}
	defer conn.Close()

	serviceClient := pb.NewMensajeServiceClient(conn)

	res, err := serviceClient.Create(context.Background(), &pb.Crearmensaje{
		Mensaje: &pb.Mensaje{
			Nombre: req.Mensaje.Nombre,
		},
	})

	if err != nil {
		panic("no se creo el mensaje" + err.Error())
	}

	fmt.Println("Estado enviado:", res.Mensajeid)

	return &pb.Respuestamensaje{
		Mensajeid: req.Mensaje.Nombre,
	}, nil
}

func (s *server) CreateLista(ctx context.Context, req *pb.ConsultarLista) (*pb.RespuestaLista, error) {

	fmt.Println("Solicitud de " + req.Estado.Nombre + " recibida")

	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())

	if err != nil {
		panic("no se puede conectar con el servidor" + err.Error())
	}
	defer conn.Close()

	serviceClient := pb.NewMensajeServiceClient(conn)

	res, err := serviceClient.CreateLista(context.Background(), &pb.ConsultarLista{
		Estado: &pb.Estado{
			Nombre: req.Estado.Nombre,
		},
	})

	if err != nil {
		panic("no se creo el mensaje" + err.Error())
	}

	return &pb.RespuestaLista{
		Estadoid: res.Estadoid,
	}, nil

}

func (s *server) CreateActualiza(ctx context.Context, req *pb.CrearActualizacion) (*pb.RespuestaActualizacion, error) {
	return nil, fmt.Errorf("CreateActualiza is not implemented")
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
