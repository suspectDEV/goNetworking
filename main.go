/*
	Main
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	UDPDaytimeClient()
}

/*	IP
 */
func typeIP() {
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

/*	Mask
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
	fmt.Println("Address is ", addr.String(),
		"Default mask length is ", bits,
		"Leading ones count is ", ones,
		"Mask is (hex) ", mask.String(),
		" Network is ", network.String())
	os.Exit(0)
}

/*	ResolveIP
 */
func ResolveIPAddr() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage; %s hostname\n", os.Args[0])
		fmt.Println("Usage: ", os.Args[0], "hostname")
		os.Exit(1)
	}

	name := os.Args[1]
	addr, err := net.ResolveIPAddr("ip", name)
	if err != nil {
		fmt.Println("Resolution error", err.Error())
		os.Exit(1)
	}
	fmt.Println("Resolved address is ", addr.String())
	os.Exit(0)
}

/*	LookupHost
 */
func LookupHost() {
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

/*	LookupPort
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

/*	GetHeadInfo
 */
func GetHeadInfo() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	result, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Println(string(result))
	os.Exit(0)
}

/*	DaytimeServer
 */
func ListenTCP() {
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
		daytime := time.Now().String()
		log.Println(daytime)
		conn.Write([]byte(daytime))
		conn.Close()
	}
}

/*	SimpleEchoServer
	..single-threaded
*/
// func SimpleEchoServer() {
// 	service := ":1201"
// 	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
// 	checkError(err)

// 	listener, err := net.ListenTCP("tcp", tcpAddr)
// 	checkError(err)

// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			continue
// 		}
// 		handleClient(conn)
// 		conn.Close()
// 	}
// }

// func handleClient(conn net.Conn) {
// 	var buf [512]byte
// 	for {
// 		n, err := conn.Read(buf[0:])
// 		if err != nil {
// 			return
// 		}
// 		fmt.Println(string(buf[0:]))
// 		_, err2 := conn.Write(buf[0:n])
// 		if err2 != nil {
// 			return
// 		}
// 	}
// }
/*	END SimpleEchoServer
	..single-threaded
*/

/*	ThreadedEchoServer
 */
// func ThreadedEchoServer() {
// 	service := ":1201"
// 	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
// 	checkError(err)

// 	listener, err := net.ListenTCP("tcp", tcpAddr)
// 	checkError(err)
// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			continue
// 		}
// 		go handleClient(conn)
// 	}
// }

// func handleClient(conn net.Conn) {
// 	defer conn.Close()

// 	var buf [512]byte
// 	for {
// 		n, err := conn.Read(buf[0:])
// 		if err != nil {
// 			return
// 		}

// 		_, err2 := conn.Write(buf[0:n])
// 		if err2 != nil {
// 			return
// 		}
// 	}
// }
/*	END ThreadedEchoServer
 */

/* ---------------------------------------------------------------
	*	3.8 UDP Datagrams:
--------------------------------------------------------------- */
/*	UDPDaytimeClient
 */
func UDPDaytimeClient() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)

	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)

	_, err = conn.Write([]byte("anything"))
	checkError(err)

	var buf [512]byte
	n, err := conn.Read(buf[0:])
	checkError(err)

	fmt.Println(string(buf[0:n]))
	os.Exit(0)
}

/*
	END UDPDaytimeClient */

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s ", err.Error())
		os.Exit(1)
	}
}
