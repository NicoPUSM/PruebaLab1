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
	fmt.Println("Solicitud de " + req.Mensaje.Nombre + " recibida, mensaje enviado: " + req.Mensaje.Nombre)

	palabras := strings.Split(req.Mensaje.Nombre, " ")

	archivo, err := os.OpenFile(palabras[1]+".txt", os.O_RDWR|os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error al abrir el archivo", err)
		return nil, err
	}

	log_registro, err := os.OpenFile("Log_Registro.txt", os.O_RDWR|os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error al abrir el archivo", err)
		return nil, err
	}

	defer archivo.Close()

	defer log_registro.Close()

	if palabras[0] == "AgregarBase" {
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

	listner, err := net.Listen("tcp", ":50052")

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
