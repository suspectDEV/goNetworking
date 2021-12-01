package main

import (
	"bytes"
	"encoding/asn1"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	if os.Args[1] == "server"{
		ASN1DaytimeServer()
	}else{
		ASN1DaytimeClient()
	}
}

/*	ASN.1
	----
	Convertir un dato (string en este caso) en arreglo de bytes (marshal)
	y luego volverlo a su estado normal (unmarshal)
*/
func ASN1() {
	mdata, err := asn1.Marshal("13")
	fmt.Println(mdata)
	checkError(err)

	var n string
	_, err1 := asn1.Unmarshal(mdata, &n)
	checkError(err1)

	fmt.Println("After marshal/unmarshal: ", n)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}


/*	ASN1 DAYTIME SERVER
	----
	Servicio en el puerto 1200
	Marshal de la hora del servidor
*/
func ASN1DaytimeServer() {
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		daytime := time.Now()
		mdata, _ := asn1.Marshal(daytime)
		conn.Write(mdata)
		conn.Close()
	}
}


/* ASN1 DAYTIME CLIENT
	----
	Cliente conectado al servidor (puerto 1200)
	UnMarshal de la respuesta del servidor
*/
func ASN1DaytimeClient(){
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage %s host:port", os.Args[0])
	}
	service := os.Args[2]
	conn, err := net.Dial("tcp", service)
	checkError(err)

	result, err := readFully(conn)
	checkError(err)

	var newtime time.Time
	_, err1 := asn1.Unmarshal(result, &newtime)
	checkError(err1)

	fmt.Println("After marshal/unmarshal: ", newtime.String())
	os.Exit(0)
}

/*
	Complemento de la función anterior
	Lee los datos que envía el servidor.
*/
func readFully(conn net.Conn) ([]byte, error){
	defer conn.Close()
	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for{
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil{
			if err == io.EOF{
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}