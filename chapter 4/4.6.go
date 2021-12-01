package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
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
	// Descomentar para guardar archivo GOB
	// ...
	// person := Person{
	// 	Name: Name{Family: "Newmarch", Personal: "Jan"},
	// 	Email: []Email{{Kind: "home", Address: "jan@newmarch.name"},
	// 		{Kind: "work", Address: "j.newmarch@boxhill.edu.au"}}}
	// saveGob("person.gob", person)
	//
	// Descomentar para cargar archivo GOB
	// ...
	// var person Person
	// LoadGob("person.gob", &person)
	// fmt.Println("Person", person.String())
	//
	if os.Args[1] == "server" {
		GobEchoServer()
	} else {
		GobEchoClient()
	}
}

/*	SAVE GOB
	----
	Se crea y almacena un archivo con extensión gob
	(propio de GO)
*/
func saveGob(filename string, key interface{}) {
	outFile, err := os.Create(filename)
	checkError(err)
	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
	outFile.Close()
}

/*	LOAD GOB
	----
	Lee un archivo .gob previamente creado
*/
func LoadGob(filename string, key interface{}) {
	infile, err := os.Open(filename)
	checkError(err)
	decoder := gob.NewDecoder(infile)
	err = decoder.Decode(key)
	checkError(err)
	infile.Close()
}

/*
	Complemento de la función anterior para hacer
	lectura	del archivo Gob
*/
func (p Person) String() string {
	s := p.Name.Personal + " " + p.Name.Family
	for _, v := range p.Email {
		s += "\n" + v.Kind + ": " + v.Address
	}
	return s
}

/*	GOB ECHO CLIENT
	----
*/
func GobEchoClient() {
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

	encoder := gob.NewEncoder(conn)
	decoder := gob.NewDecoder(conn)

	for n := 0; n < 10; n++ {
		encoder.Encode(person)
		var newPerson Person
		decoder.Decode(&newPerson)
		fmt.Println(newPerson.String())
	}
	os.Exit(0)
}

/*
	Complemento de la función anterior para hacer lectura de datos
	en un buffer creado
	// TODO: Esta función no se está ejecutando (revisar)
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

/*	GOB ECHO SERVER
	----
*/
func GobEchoServer() {
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
		encoder := gob.NewEncoder(conn)
		decoder := gob.NewDecoder(conn)

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
		panic(err)
	}
}
