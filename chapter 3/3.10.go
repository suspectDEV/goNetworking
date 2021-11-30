package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	ThreadedIPEchoServer()
}

/*	IP GET HEAD INFO
	----
	Reescribiendo la función del punto 3.6
	Traer la información de la cabecera de un sitio usando una función general
	Solo "Dial" en cambio de "DialTCP"
*/
func IPGetHeadInfo() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	conn, err := net.Dial("tcp", service)
	checkError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	result, err := readFully(conn)
	checkError(err)

	fmt.Println(string(result))
	os.Exit(0)
}

/*
	----
	Esta función actúa como complemento de la función anterior
	Se encarga de hacer lectura en bytes de la conexión
*/
func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}


/*	THREADED IP ECHO SERVER
	----
	Crea servicio en el puerto 1200 haciendo un echo
	de lo que recibe del cliente
*/
func ThreadedIPEchoServer(){
	service := ":1200"
	listener, err := net.Listen("tcp", service)
	checkError(err)

	for{
		conn, err := listener.Accept()
		if err != nil{
			continue
		}
		go handleClient(conn)
	}
}

/*
	Complemento de la función anterior crea bytes para leer y escribir
	respuesta al cliente
*/
func handleClient(conn net.Conn){
	defer conn.Close()
	var buf [512]byte
	for{
		n, err := conn.Read(buf[0:])
		if err != nil{
			return
		}
		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
