package main

import (
	"chatroom/server/model"
	"fmt"
	"net"
	"time"
)

func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {
	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()
	fmt.Println("The server (new structure) is listening on port 8889 ...")
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
		go processMain(conn)
	}
}

func processMain(conn net.Conn) {
	defer conn.Close()
	processor := &Processor{Conn: conn}
	err := processor.mainProcess()
	if err != nil {
		fmt.Println("Client and server communication goroutine error:", err)
		return
	}
}
