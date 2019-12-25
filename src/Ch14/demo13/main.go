package main

import (
	"flag"
	"fmt"
)

func main() {
	var user string
	var pwd string
	var host string
	var port int

	flag.StringVar(&user, "u", "", "User name, default is blank.")
	flag.StringVar(&pwd, "pwd", "", "Password, default is blank.")
	flag.StringVar(&host, "h", "", "Host name, default is blank.")
	flag.IntVar(&port, "port", 3306, "Port number.")

	flag.Parse()

	fmt.Println("user =", user)
	fmt.Println("pwd =", pwd)
	fmt.Println("host =", host)
	fmt.Println("port =", port)
}
