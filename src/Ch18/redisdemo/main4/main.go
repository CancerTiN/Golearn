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

	r1, err := conn.Do("HMSET", "user02", "name", "john", "age", 19)
	if err != nil {
		fmt.Println("conn.Do(HMSET) error:", err)
		return
	}
	fmt.Println("conn.Do(HMSET) success:", r1)

	r2, err := redis.Strings(conn.Do("HMGET", "user01", "name", "age"))
	if err != nil {
		fmt.Println("conn.Do(HMGET) error:", err)
		return
	}
	fmt.Println("conn.Do(HMGET) success:", r2)
	fmt.Printf("The type of r2 is %T\n", r2)
	fmt.Printf("The value of r2 is %v\n", r2)

	for i, v := range r2 {
		fmt.Printf("r2[%v] = %v\n", i, v)
	}
}
