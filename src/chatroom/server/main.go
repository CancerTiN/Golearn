package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		fmt.Println("Read the data sent by the client ...")
		buf := make([]byte, 8096)
		n, err := conn.Read(buf[:4])
		if err != nil || n != 4 {
			fmt.Println("conn.Read error:", err)
			return
		}
		fmt.Println("Received data from the client:", buf[:4])
	}
}

func main() {
	fmt.Println("The server is listening on port 8889 ...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Listen error:", err)
		return
	}
	defer listen.Close()
	fmt.Println("The server is waiting for the client to connect ...")
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept error:", err)
		}
		go process(conn)
	}
}
