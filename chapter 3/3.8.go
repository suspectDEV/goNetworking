package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

// EX 3.8: UDP Datagrams
func main(){
	if os.Args[1] == "server"{
		UDPtimeServer()
	}else{
		UDPDaytimeClient()
	}
}

/*	UDP TIME SERVER
	----
	Crea una conexión de servidor en el puerto 1200
	*/
func UDPtimeServer(){
	service := ":1200"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)

	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)

	for{
		handleClient(conn)
	}
}

/*	HANDLE CLIENT
	----
	Se ejecuta cuando un cliente se conecta al servidor
	Da como respuesta la fecha actual del servidor
	*/
func handleClient(conn *net.UDPConn){
	var buf [512]byte
	_, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil{
		return
	}
	daytime := time.Now().String()
	conn.WriteToUDP([]byte(daytime), addr)
}


/*	UDP DAYTIME CLIENT
	----
	Cliente que se conecta al servidor, especificar el puerto 1200
	Se crea un buffer de 512 bytes que recibe la información del servidor.
	*/
func UDPDaytimeClient(){
	if len(os.Args) != 3{
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[2]
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)

	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)

	_, err = conn.Write([]byte("anything"))
	checkError(err)

	var buf [512]byte
	n, err := conn.Read(buf[0:])
	checkError(err)

	fmt.Println(string(buf[0:n]))
	os.Exit(0)
}

func checkError(err error){
	if err != nil{
		panic(err)
	}
}