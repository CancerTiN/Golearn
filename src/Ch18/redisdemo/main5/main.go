package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
	}
}

func main() {
	defer pool.Close()
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", "name", "Tom Cat ~")
	if err != nil {
		fmt.Println("conn.Do(SET) error:", err)
		return
	}

	r, err := conn.Do("GET", "name")
	if err != nil {
		fmt.Println("conn.Do(GET) error:", err)
		return
	}
	fmt.Println("conn.Do(GET) success:", r)
}
