package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"
)

// EX 3.6: TCP Sockets
// TODO: Falta probar SimpleEchoServer / ThreadedEchoServer
func main() {
	ThreadedEchoServer()
}

/* GETHEADINFO
----
Trae la información de la cabecera de un sitio
TODO: Actualmente con grubbe.io - tiza.la no trae info con el puerto 80
*/
func GetHeadInfo() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	result, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Println(string(result))
	os.Exit(0)
}

/* DAYTIMESERVER
----
Crea un servicio en el puerto 1200: Responde la hora a los clientes
que se conecten
*/
func DaytimeServer() {
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Continuo")
			continue
		}
		daytime := time.Now().String()
		log.Println(daytime)
		conn.Write([]byte(daytime))
		conn.Close()
	}
}

/* SIMPLE ECHO SERVER
----
Levanta un servicio de escucha en el puerto :1201
*/
// func SimpleEchoServer() {
// 	service := ":1201"
// 	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
// 	checkError(err)

// 	listener, err := net.ListenTCP("tcp", tcpAddr)
// 	checkError(err)

// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			continue
// 		}
// 		handleClient(conn)
// 		conn.Close()
// 	}
// }

// func handleClient(conn net.Conn) {
// 	var buf [512]byte
// 	for {
// 		n, err := conn.Read(buf[0:])
// 		if err != nil {
// 			return
// 		}
// 		_, err2 := conn.Write(buf[0:n])
// 		if err2 != nil {
// 			return
// 		}
// 	}
// }

/*	THREADED ECHO SERVER
 */
func ThreadedEchoServer() {
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
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
