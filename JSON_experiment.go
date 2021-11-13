// package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net"
// 	"os"
// )

// func main() {
// 	if os.Args[1] == "server" {
// 		EchoServer()
// 	} else {
// 		EchoClient()
// 	}
// }

// /* JSON EchoClient
//  */
// type Person struct {
// 	Name  Name
// 	Email []Email
// }

// type Name struct {
// 	Family   string
// 	Personal string
// }

// type Email struct {
// 	Kind    string
// 	Address string
// }

// func (p Person) String() string {
// 	s := p.Name.Personal + " " + p.Name.Family
// 	for _, v := range p.Email {
// 		s += "\n" + v.Kind + ": " + v.Address
// 	}
// 	return s
// }

// func EchoClient() {
// 	person := Person{
// 		Name: Name{Family: "Newmarch", Personal: "Jan"},
// 		Email: []Email{{Kind: "home", Address: "jan@newmarch.name"},
// 			{Kind: "work", Address: "j.newmarch@boxhill.edu.au"}}}

// 	if len(os.Args) != 3 {
// 		fmt.Println("Usage: ", os.Args[0], "host:port")
// 		os.Exit(1)
// 	}
// 	service := os.Args[2]
// 	conn, err := net.Dial("tcp", service)
// 	checkError(err)

// 	encoder := json.NewEncoder(conn)
// 	decoder := json.NewDecoder(conn)

// 	for n := 0; n < 10; n++ {
// 		encoder.Encode(person)
// 		var newPerson Person
// 		decoder.Decode(&newPerson)
// 		fmt.Println(newPerson.String())
// 	}
// 	os.Exit(0)
// }

// func readFully(conn net.Conn) ([]byte, error) {
// 	defer conn.Close()

// 	result := bytes.NewBuffer(nil)
// 	var buf [512]byte
// 	for {
// 		n, err := conn.Read(buf[0:])
// 		result.Write(buf[0:n])
// 		if err != nil {
// 			if err == io.EOF {
// 				break
// 			}
// 			return nil, err
// 		}
// 	}
// 	return result.Bytes(), nil
// }

// /* END JSON EchoClient
//  */

// /*	JSON EchoServer
//  */
// func EchoServer() {
// 	service := "0.0.0.0:1200"
// 	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
// 	checkError(err)

// 	listener, err := net.ListenTCP("tcp", tcpAddr)
// 	checkError(err)

// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			continue
// 		}
// 		encoder := json.NewEncoder(conn)
// 		decoder := json.NewDecoder(conn)

// 		for n := 0; n < 10; n++ {
// 			var person Person
// 			decoder.Decode(&person)
// 			fmt.Println(person.String())
// 			encoder.Encode(person)
// 		}
// 		conn.Close()
// 	}
// }

// /*	END JSON EchoServer
//  */

// func checkError(err error) {
// 	if err != nil {
// 		fmt.Println("Fatal error ", err.Error())
// 		os.Exit(1)
// 	}
// }
