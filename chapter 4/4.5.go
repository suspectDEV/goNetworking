package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type Person struct {
	Name  Name
	Email []Email
}

type Name struct {
	Family   string
	Personal string
}

type Email struct {
	Kind    string
	Address string
}

func main() {
	// Descomentar para usar función SaveJSON
	// ..
	// person := Person{
	// 	Name: Name{Family: "Forero", Personal: "Alexander"},
	// 	Email: []Email{{Kind: "personal", Address: "alex@grubbe.io"},
	// 		{Kind: "work", Address: "support@tiza.la"}}}
	// SaveJSON("person.json", person)
	// ..
	// Descomentar para usar función LoadJSON
	// var person Person
	// LoadJSON("person.json", &person)
	// fmt.Println("Person ", person.String())
	if os.Args[1] == "server" {
		JSONEchoServer()
	} else {
		JSONEchoClient()
	}
}

/*	SAVE JSON
	----
	Crea un archivo con extensión JSON
	y almacena los datos allí.
*/
func SaveJSON(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	checkError(err)
	encoder := json.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
	outFile.Close()
}

/* LOAD JSON
----
Carga un archivo JSON previamente
creado
*/
func LoadJSON(filename string, key interface{}) {
	inFile, err := os.Open(filename)
	checkError(err)
	decoder := json.NewDecoder(inFile)
	err = decoder.Decode(key)
	checkError(err)
	inFile.Close()
}

/*
	Complemento de la función anterior para traer
	datos del archivo JSON en un lenguaje legible
*/
func (p Person) String() string {
	s := p.Name.Personal + " " + p.Name.Family
	for _, v := range p.Email {
		s += "\n" + v.Kind + ": " + v.Address
	}
	return s
}

/*	JSON ECHO CLIENT
	----
	Cliente conectado al servicio del puerto 1200 que envía los datos
	de un archivo JSON con 10 lecturas de ciclo for
*/
func JSONEchoClient() {
	person := Person{
		Name: Name{Family: "Newmarch", Personal: "Jan"},
		Email: []Email{{Kind: "home", Address: "jan@newmarch.name"},
			{Kind: "work", Address: "j.newmarch@boxhill.edu.au"}}}

	if len(os.Args) != 3 {
		fmt.Println("Usage: ", os.Args[0], "host:port")
		os.Exit(1)
	}
	service := os.Args[2]

	conn, err := net.Dial("tcp", service)
	checkError(err)

	encoder := json.NewEncoder(conn)
	decoder := json.NewDecoder(conn)

	for n := 0; n < 10; n++ {
		encoder.Encode(person)
		var newPerson Person
		decoder.Decode(&newPerson)
		fmt.Println(newPerson.String())
	}
	os.Exit(0)
}

/*	JSON ECHO SERVER
	----
	Crea un servicio en el puerto 1200 capaz de leer y escribir
	un archivo JSON 10 veces enviado por el cliente.
*/
func JSONEchoServer() {
	service := "0.0.0.0:1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		encoder := json.NewEncoder(conn)
		decoder := json.NewDecoder(conn)

		for n := 0; n < 10; n++ {
			var person Person
			decoder.Decode(&person)
			fmt.Println(person.String())
			encoder.Encode(person)
		}
		conn.Close()
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
