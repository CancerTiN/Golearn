package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379", Password: "", DB: 0})
	pong, err := rdb.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("rdb.Ping().Result() return: (%#v, %#v).\n", pong, err)
}

func demoAlpha() {
	key := "rank"
	members := []*redis.Z{
		&redis.Z{Score: 90, Member: "PHP"},
		&redis.Z{Score: 96, Member: "Golang"},
		&redis.Z{Score: 97, Member: "Python"},
		&redis.Z{Score: 99, Member: "Java"},
	}
	intCmd := rdb.ZAdd(key, members...)
	fmt.Printf("rdb.ZAdd(%#v, %#v...) return: (%#v).\n", key, members, intCmd)
}

func demoBeta() {
	key := "rank"
	increment := 10.0
	member := "Golang"
	floatCmd := rdb.ZIncrBy(key, increment, member)
	fmt.Printf("rdb.ZIncrBy(%#v, %#v, %#v) return: (%#v).\n", key, increment, member, floatCmd)
}

func main() {
	fmt.Println(rdb)
	switch 2 {
	case 1:
		demoAlpha()
	case 2:
		demoBeta()
	}
}
