package main

import (
	"fmt"
	"net"
	"os"
)

// EX 3.4: IP address type (p. 15-18)
// ----------
func main() {
	Lookuphost()
}

/*	IP
	----
	Convierte un string enviado en un dirección IP según corresponda
	IPV4 IPV6
*/
func IP() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	}
	os.Exit(0)
}

/*	IP MASK
	----
	Trae datos de la IP ingresada (bits, ones, mask_hex)
*/
func IPMask() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s dotted-ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	dotAddr := os.Args[1]
	addr := net.ParseIP(dotAddr)
	if addr == nil {
		fmt.Println("Invalid address")
		os.Exit(1)
	}
	mask := addr.DefaultMask()
	network := addr.Mask(mask)
	ones, bits := mask.Size()
	fmt.Println(
		"Address is ", addr.String(),
		"\nDefault mask length is ", bits,
		"\nLeading ones count is ", ones,
		"\nMask is (hex) ", mask.String(),
		"\nNetwork is ", network.String())
	os.Exit(0)
}

/* RESOLVE IP
----
Ingresa nombre de dominio y devuelve la IP
*/
func ResolveIPAddr() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage; %s hostname\n", os.Args[0])
		fmt.Println("Usage: ", os.Args[0], "hostname")
		os.Exit(1)
	}

	domain := os.Args[1]
	addr, err := net.ResolveIPAddr("ip", domain)
	if err != nil {
		fmt.Println("Resolution error", err.Error())
		os.Exit(1)
	}
	fmt.Println("Resolved address is ", addr.String())
	os.Exit(0)
}

/* LOOKUPHOST
----
Ingresa dominio/IP y devuelve las IP que le pertenecen
Trae varios datos generalmente cuando se explora una "nube" (aws, fb, google)
*/
func Lookuphost() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	addrs, err := net.LookupHost(name)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(2)
	}
	for _, s := range addrs {
		fmt.Println(s)
	}
	os.Exit(0)
}
