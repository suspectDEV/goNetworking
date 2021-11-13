package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	LoadJSONf()
}

/* ANS1 DaytimeServer
 */
// func DaytimeServer() {
// 	service := ":1200"
// 	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
// 	checkError(err)
// 	listener, err := net.ListenTCP("tcp", tcpAddr)
// 	checkError(err)

// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			continue
// 		}

// 		daytime := time.Now()
// 		mdata, _ := asn1.Marshal(daytime)
// 		conn.Write(mdata)
// 		conn.Close()
// 	}
// }

/* END ANS1 DaytimeServer
 */

/*	ASN.1 DaytimeClient
 */
// func DaytimeClient() {
// 	if len(os.Args) != 3 {
// 		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
// 		os.Exit(1)
// 	}
// 	service := os.Args[2]
// 	conn, err := net.Dial("tcp", service)
// 	checkError(err)

// 	result, err := readFully(conn)
// 	checkError(err)

// 	var newtime time.Time
// 	_, err1 := asn1.Unmarshal(result, &newtime)
// 	checkError(err1)

// 	fmt.Println("After marshal/unmarshal: ", newtime.String())
// 	os.Exit(0)
// }
/*	END ASN.1 DaytimeClient
 */

/* Save JSON
 */

type DATA struct {
	Usuario Usuario
	Clase   Clase
}

type Usuario struct {
	Nombre string
}

type Clase struct {
	Nombre string
}

func CreateJSON() {
	usuario := DATA{
		Usuario: Usuario{Nombre: "Alexander"},
		Clase:   Clase{Nombre: os.Args[1]},
	}
	saveJSON("myJSON.json", usuario)
}

func saveJSON(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	checkError(err)
	encoder := json.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
	outFile.Close()
}

/* END Save JSON
 */

/*	LoadJSON
 */
func (D DATA) String() string {
	s := "Usuario: " + D.Usuario.Nombre + "\n"
	s += "Clase: " + D.Clase.Nombre

	return s
}

func LoadJSONf() {
	var data DATA
	loadJSON("myJSON.json", &data)
	fmt.Println("JSON_DATA: ", data)
}

func loadJSON(fileName string, key interface{}) {
	inFile, err := os.Open(fileName)
	checkError(err)
	decoder := json.NewDecoder(inFile)
	err = decoder.Decode(key)
	checkError(err)
	inFile.Close()
}

/* END LoadJSON
 */

// func checkError(err error) {
// 	if err != nil {
// 		fmt.Fprint(os.Stderr, "Fatal error: %s", err.Error())
// 		os.Exit(1)
// 	}
// }

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
