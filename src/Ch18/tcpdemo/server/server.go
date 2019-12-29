package main

import (
	"fmt"
	"io"
	"net"
	"reflect"
)

func main() {
	fmt.Println("The server starts listening ...")
	listener, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("Failed to start listening:", err)
		return
	}
	defer listener.Close() // Delay closing the listener
	fmt.Println("Start listening successfully, listener is", listener)
	fmt.Println("Wait for the client to connect ...")
	for {
		conn, err := listener.Accept() // Wait for the client to connect
		if err != nil {
			fmt.Println("Failed to accept waits for:", err)
		} else {
			fmt.Println("Client successfully connected, connection is", conn)
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("The server is waiting for the client (%v) to send information ...\n", conn.RemoteAddr().String())
	for {
		buf := make([]byte, 1024)
		// Wait for the client to send information through the connection.
		// If the client does not send, the goroutine is blocked here.
		n, err := conn.Read(buf)
		if err == io.EOF {
			fmt.Println("Client has exited:", err)
		} else if reflect.TypeOf(err) == reflect.TypeOf(&net.OpError{}) {
			fmt.Println("An exception (net.OpError) occurs:", err)
			return
		} else if err != nil {
			fmt.Println("Read data from the connection error:", err)
			fmt.Println("The type of error is", reflect.TypeOf(err))
			return
		}
		fmt.Println("Received client information:")
		fmt.Print(string(buf[:n])) // !!!
	}
}
