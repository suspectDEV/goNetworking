package main

import (
	"fmt"
	"net"
	"os"
)

// EX 3.5: Services (p. 18)
func main() {
	LookupPort()
}

/*	LOOKUPPORT
	----
	Esta función permite identificar el puerto que está utilizando un servicio
*/
func LookupPort() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s network-type service\n", os.Args[0])
		os.Exit(1)
	}
	networkType := os.Args[1]
	service := os.Args[2]
	port, err := net.LookupPort(networkType, service)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(2)
	}
	fmt.Println("Service port ", port)
	os.Exit(0)
}
