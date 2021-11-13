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
	if os.Args[1] == "server" {
		DaytimeServer()
	} else {
		DaytimeClient()
	}
}

/* ANS1 DaytimeServer
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
			continue
		}

		daytime := time.Now()
		mdata, _ := asn1.Marshal(daytime)
		conn.Write(mdata)
		conn.Close()
	}
}

/* END ANS1 DaytimeServer
 */

/*	ASN.1 DaytimeClient
 */
func DaytimeClient() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
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

/*	END ASN.1 DaytimeClient
 */

func checkError(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

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
