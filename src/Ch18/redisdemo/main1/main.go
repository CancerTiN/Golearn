package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial error:", err)
		return
	}
	fmt.Println("redis.Dial success:", conn)
	defer conn.Close()

	r1, err := conn.Do("SET", "name", "tom|jerry")
	if err != nil {
		fmt.Println("conn.Do(SET) error:", err)
		return
	}
	fmt.Println("conn.Do(Set) success:", r1)

	r2, err := conn.Do("GET", "name")
	if err != nil {
		fmt.Println("conn.Do(GET) error:", err)
		return
	}
	fmt.Println("conn.Do(GET) success:", r2)
	fmt.Printf("The type of r2 is %T\n", r2)

	r3, err := redis.String(conn.Do("GET", "name"))
	if err != nil {
		fmt.Println("conn.Do(GET) error:", err)
		return
	}
	fmt.Println("conn.Do(GET) success:", r3)
	fmt.Printf("The type of r3 is %T\n", r3)
	fmt.Printf("The value of r3 is %v\n", r3)
}
