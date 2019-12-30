package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.56.1:8888")
	if err != nil {
		fmt.Println("Failed to dial connects:", err)
	} else {
		fmt.Println("Successfully connected to the server, connection is", conn)
		fmt.Println("The remote address is", conn.RemoteAddr().String())
	}
	var nTotalBytes int
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("The client is waiting for user input:")
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Read string from stdin error:", err)
		}
		if strings.Trim(str, "\r\n") == "exit" {
			fmt.Println("Client exits ...")
			break
		}
		n, err := conn.Write([]byte(str))
		if err != nil {
			fmt.Println("Write data to the connection error:", err)
		}
		nTotalBytes += n
	}
	fmt.Printf("The client sent %d bytes of data and exited", nTotalBytes)
}
