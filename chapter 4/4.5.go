package main

import (
	"encoding/json"
	"fmt"
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
	var person Person
	LoadJSON("person.json", &person)
	fmt.Println("Person ", person.String())
}

/*	SaveJSON
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

/* LoadJSON
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

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
