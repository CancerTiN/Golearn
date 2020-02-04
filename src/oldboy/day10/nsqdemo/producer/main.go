package main

import (
	"bufio"
	"fmt"
	"github.com/nsqio/go-nsq"
	"os"
	"strings"
)

var producer *nsq.Producer

func init() {
	var (
		addr   = "127.0.0.1:4150"
		config = nsq.NewConfig()
		err    error
	)
	producer, err = nsq.NewProducer(addr, config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("nsq.NewProducer(%#v, %#v) return: (%#v, %#v).\n", addr, config, producer, err)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Message to be published: ")
		s, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("reader.ReadString('\n') return: (%#v, %#v).\n", s, err)
			continue
		}
		s = strings.TrimSpace(s)
		if strings.ToUpper(s) == "Q" {
			break
		}
		topic := "topic_demo"
		body := []byte(s)
		err = producer.Publish(topic, body)
		if err != nil {
			fmt.Printf("producer.Publish(%#v, %#) return: (%#v).\n", topic, body, err)
			continue
		}
	}
}
