package main

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	pb "github.com/NicoPUSM/PruebaLab1/Prueba-Lab1/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnsafeMensajeServiceServer
	relojVectorial []int
}

func (s *server) Create(ctx context.Context, req *pb.Crearmensaje) (*pb.Respuestamensaje, error) {
	//Extra
	conn1, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		panic("no se puede conectar con el servidor 1" + err.Error())
	}
	defer conn1.Close()

	serviceClient1 := pb.NewMensajeServiceClient(conn1)

	_, err := serviceClient1.CreateActualiza(context.Background(), &pb.CrearActualizacion{
		Actualiza: &pb.Actualizar{
			Nombre: req.Mensaje.Nombre,
		},
	})

	if err != nil {
		panic("no se creo el mensaje en el servidor 1" + err.Error())
	}

	conn2, err := grpc.Dial("localhost:50054", grpc.WithInsecure())
	if err != nil {
		panic("no se puede conectar con el servidor 2" + err.Error())
	}
	defer conn2.Close()

	serviceClient2 := pb.NewMensajeServiceClient(conn2)

	_, err := serviceClient2.CreateActualiza(context.Background(), &pb.CrearActualizacion{
		Actualiza: &pb.Actualizar{
			Nombre: req.Mensaje.Nombre,
		},
	})

	if err != nil {
		panic("no se creo el mensaje en el servidor 2" + err.Error())
	}

	//Importante

	fmt.Println("Solicitud de " + req.Mensaje.Nombre + " recibida, mensaje enviado: " + req.Mensaje.Nombre)

	palabras := strings.Split(req.Mensaje.Nombre, " ")

	log_registro, err := os.OpenFile("Log_Registro.txt", os.O_RDWR|os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error al abrir el archivo", err)
		return nil, err
	}

	defer log_registro.Close()

	if palabras[0] == "AgregarBase" {
		archivo, err := os.OpenFile(palabras[1]+".txt", os.O_RDWR|os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("Error al abrir el archivo", err)
			return nil, err
		}

		defer archivo.Close()

		_, err = archivo.WriteString(palabras[2] + " " + palabras[3] + "\n")

		if err != nil {
			fmt.Println("Error al escribir en el archivo", err)
			return nil, err
		}

		_, err = log_registro.WriteString(palabras[0] + " " + palabras[1] + " " + palabras[2] + " " + palabras[3] + "\n")

		if err != nil {
			fmt.Println("Error al escribir en el archivo", err)
			return nil, err
		}

	} else if palabras[0] == "RenombrarBase" {
		archivo, err := os.OpenFile(palabras[1]+".txt", os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("Error al abrir el archivo", err)
			return nil, err
		}

		defer archivo.Close()

		var nuevoContenido []string
		scanner := bufio.NewScanner(archivo)

		for scanner.Scan() {
			line := scanner.Text()
			div := strings.Fields(line)

			if div[0] == palabras[2] {
				div[0] = palabras[3]
				line = strings.Join(div, " ")
			}

			nuevoContenido = append(nuevoContenido, line)
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error al escanear el archivo:", err)
			return nil, err
		}

		if _, err := archivo.Seek(0, 0); err != nil {
			fmt.Println("Error al rebobinar el archivo:", err)
			return nil, err
		}

		if err := archivo.Truncate(0); err != nil {
			fmt.Println("Error al truncar el archivo:", err)
			return nil, err
		}

		for _, linea := range nuevoContenido {
			fmt.Fprintln(archivo, linea)
		}

		_, err = log_registro.WriteString(palabras[0] + " " + palabras[1] + " " + palabras[2] + " " + palabras[3] + "\n")

		if err != nil {
			fmt.Println("Error al escribir en el archivo", err)
			return nil, err
		}

	} else if palabras[0] == "ActualizarValor" {
		archivo, err := os.OpenFile(palabras[1]+".txt", os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("Error al abrir el archivo", err)
			return nil, err
		}

		defer archivo.Close()

		var nuevoContenido []string
		scanner := bufio.NewScanner(archivo)

		for scanner.Scan() {
			line := scanner.Text()
			div := strings.Fields(line)

			if div[0] == palabras[2] {
				div[1] = palabras[3]
				line = strings.Join(div, " ")
			}

			nuevoContenido = append(nuevoContenido, line)
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error al escanear el archivo:", err)
			return nil, err
		}

		if _, err := archivo.Seek(0, 0); err != nil {
			fmt.Println("Error al rebobinar el archivo:", err)
			return nil, err
		}

		if err := archivo.Truncate(0); err != nil {
			fmt.Println("Error al truncar el archivo:", err)
			return nil, err
		}

		for _, linea := range nuevoContenido {
			fmt.Fprintln(archivo, linea)
		}

		_, err = log_registro.WriteString(palabras[0] + " " + palabras[1] + " " + palabras[2] + " " + palabras[3] + "\n")

		if err != nil {
			fmt.Println("Error al escribir en el archivo", err)
			return nil, err
		}

	} else if palabras[0] == "BorrarBase" {
		archivo, err := os.OpenFile(palabras[1]+".txt", os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("Error al abrir el archivo", err)
			return nil, err
		}

		defer archivo.Close()

		var nuevoContenido []string
		scanner := bufio.NewScanner(archivo)

		for scanner.Scan() {
			line := scanner.Text()
			div := strings.Fields(line)

			if div[0] != palabras[2] {
				nuevoContenido = append(nuevoContenido, line)
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error al escanear el archivo:", err)
			return nil, err
		}

		if err := archivo.Truncate(0); err != nil {
			fmt.Println("Error al truncar el archivo:", err)
			return nil, err
		}

		if _, err := archivo.Seek(0, 0); err != nil {
			fmt.Println("Error al rebobinar el archivo:", err)
			return nil, err
		}

		for _, linea := range nuevoContenido {
			fmt.Fprintln(archivo, linea)
		}

		_, err = log_registro.WriteString(palabras[0] + " " + palabras[1] + " " + palabras[2] + "\n")

		if err != nil {
			fmt.Println("Error al escribir en el archivo", err)
			return nil, err
		}

	}
	s.relojVectorial[0]++
	fmt.Println(s.relojVectorial)

	return &pb.Respuestamensaje{
		Mensajeid: req.Mensaje.Nombre,
	}, nil
}

func (s *server) CreateLista(ctx context.Context, req *pb.ConsultarLista) (*pb.RespuestaLista, error) {
	lista := []string{}

	archivo, err := os.OpenFile(req.Estado.Nombre+".txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error al abrir el archivo", err)
		return nil, err
	}
	defer archivo.Close()

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		line := scanner.Text()

		lista = append(lista, line)

	}

	return &pb.RespuestaLista{
		Estadoid: lista,
	}, nil

}

func (s *server) CreateActualiza(ctx context.Context, req *pb.CrearActualizacion) (*pb.RespuestaActualizacion, error) {
	fmt.Println("Solicitud de " + req.Actualiza.Nombre + " recibida, mensaje enviado: " + req.Actualiza.Nombre)

	palabras := strings.Split(req.Actualiza.Nombre, " ")

	if palabras[0] == "AgregarBase" {
		archivo, err := os.OpenFile(palabras[1]+".txt", os.O_RDWR|os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("Error al abrir el archivo", err)
			return nil, err
		}

		defer archivo.Close()

		_, err = archivo.WriteString(palabras[2] + " " + palabras[3] + "\n")

		if err != nil {
			fmt.Println("Error al escribir en el archivo", err)
			return nil, err
		}

	} else if palabras[0] == "RenombrarBase" {
		archivo, err := os.OpenFile(palabras[1]+".txt", os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("Error al abrir el archivo", err)
			return nil, err
		}

		defer archivo.Close()

		var nuevoContenido []string
		scanner := bufio.NewScanner(archivo)

		for scanner.Scan() {
			line := scanner.Text()
			div := strings.Fields(line)

			if div[0] == palabras[2] {
				div[0] = palabras[3]
				line = strings.Join(div, " ")
			}

			nuevoContenido = append(nuevoContenido, line)
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error al escanear el archivo:", err)
			return nil, err
		}

		if _, err := archivo.Seek(0, 0); err != nil {
			fmt.Println("Error al rebobinar el archivo:", err)
			return nil, err
		}

		if err := archivo.Truncate(0); err != nil {
			fmt.Println("Error al truncar el archivo:", err)
			return nil, err
		}

		for _, linea := range nuevoContenido {
			fmt.Fprintln(archivo, linea)
		}

	} else if palabras[0] == "ActualizarValor" {
		archivo, err := os.OpenFile(palabras[1]+".txt", os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("Error al abrir el archivo", err)
			return nil, err
		}

		defer archivo.Close()

		var nuevoContenido []string
		scanner := bufio.NewScanner(archivo)

		for scanner.Scan() {
			line := scanner.Text()
			div := strings.Fields(line)

			if div[0] == palabras[2] {
				div[1] = palabras[3]
				line = strings.Join(div, " ")
			}

			nuevoContenido = append(nuevoContenido, line)
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error al escanear el archivo:", err)
			return nil, err
		}

		if _, err := archivo.Seek(0, 0); err != nil {
			fmt.Println("Error al rebobinar el archivo:", err)
			return nil, err
		}

		if err := archivo.Truncate(0); err != nil {
			fmt.Println("Error al truncar el archivo:", err)
			return nil, err
		}

		for _, linea := range nuevoContenido {
			fmt.Fprintln(archivo, linea)
		}

	} else if palabras[0] == "BorrarBase" {
		archivo, err := os.OpenFile(palabras[1]+".txt", os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("Error al abrir el archivo", err)
			return nil, err
		}

		defer archivo.Close()

		var nuevoContenido []string
		scanner := bufio.NewScanner(archivo)

		for scanner.Scan() {
			line := scanner.Text()
			div := strings.Fields(line)

			if div[0] != palabras[2] {
				nuevoContenido = append(nuevoContenido, line)
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error al escanear el archivo:", err)
			return nil, err
		}

		if err := archivo.Truncate(0); err != nil {
			fmt.Println("Error al truncar el archivo:", err)
			return nil, err
		}

		if _, err := archivo.Seek(0, 0); err != nil {
			fmt.Println("Error al rebobinar el archivo:", err)
			return nil, err
		}

		for _, linea := range nuevoContenido {
			fmt.Fprintln(archivo, linea)
		}

	}

	return &pb.RespuestaActualizacion{
		Actualizaid: req.Actualiza.Nombre,
	}, nil
}

func consistenciaEventual() {
	for {
		select {
		case <-time.After(60 * time.Second):
			fmt.Println("Hola")
		}
	}
}

func main() {

	go consistenciaEventual()

	listner, err := net.Listen("tcp", ":50053")

	if err != nil {
		panic("no se creo la conexion tcp " + err.Error())
	}

	valoresIniciales := []int{0, 0, 0} // Cambia estos valores según tus necesidades.

	serv := grpc.NewServer()

	pb.RegisterMensajeServiceServer(serv, &server{relojVectorial: valoresIniciales})

	if err = serv.Serve(listner); err != nil {
		panic("no se inicio el server " + err.Error())
	}
}
