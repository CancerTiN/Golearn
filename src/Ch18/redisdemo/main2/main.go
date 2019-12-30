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

	r1, err := conn.Do("HSET", "user01", "name", "john")
	if err != nil {
		fmt.Println("conn.Do(HSET) error:", err)
		return
	}
	fmt.Println("conn.Do(HSET) success:", r1)

	r2, err := conn.Do("HSET", "user01", "age", 18)
	if err != nil {
		fmt.Println("conn.Do(HSET) error:", err)
		return
	}
	fmt.Println("conn.Do(HSET) success:", r2)

	r3, err := redis.String(conn.Do("HGET", "user01", "name"))
	if err != nil {
		fmt.Println("conn.Do(HGET) error:", err)
		return
	}
	fmt.Println("conn.Do(HGET) success:", r3)

	r4, err := redis.String(conn.Do("HGET", "user01", "age"))
	if err != nil {
		fmt.Println("conn.Do(HGET) error:", err)
		return
	}
	fmt.Println("conn.Do(HGET) success:", r4)
}
