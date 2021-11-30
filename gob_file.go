package main

// import (
// 	"encoding/gob"
// 	"fmt"
// 	"os"
// )

// /*	SaveGob
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

// func main(){
// 	mainLoad()
// }

// func mainSave() {
// 	person := Person{
// 		Name: Name{Family: "Newmarch", Personal: "Jan"},
// 		Email: []Email{{Kind: "home", Address: "jan@newmarch.name"},
// 			{Kind: "work", Address: "j.newmarch@boxhill.edu.au"}}}

// 	saveGob("person.gob", person)
// }

// func saveGob(fileName string, key interface{}) {
// 	outFile, err := os.Create(fileName)
// 	checkError(err)
// 	encoder := gob.NewEncoder(outFile)
// 	err = encoder.Encode(key)
// 	checkError(err)
// 	outFile.Close()
// }

// /*	LoadGob
// 	*/
// func (p Person) String() string{
// 	s:= p.Name.Personal + " " + p.Name.Family
// 	for _, v := range p.Email{
// 		s += "\n" + v.Kind + ": " + v.Address
// 	}
// 	return s
// }

// func mainLoad(){
// 	var person Person
// 	loadGob("person.gob", &person)
// 	fmt.Println("Person", person.String())
// }

// func loadGob(fileName string, key interface{}){
// 	inFile, err := os.Open(fileName)
// 	checkError(err)
// 	decoder := gob.NewDecoder(inFile)
// 	err = decoder.Decode(key)
// 	checkError(err)
// 	inFile.Close()
// }

// func checkError(err error) {
// 	if err != nil {
// 		fmt.Println("Fatal error ", err.Error())
// 		os.Exit(1)
// 	}
// }
