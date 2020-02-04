package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var consumer *nsq.Consumer

type MyHandler struct {
	Title string
}

func (m *MyHandler) HandleMessage(message *nsq.Message) (err error) {
	fmt.Println("m =", m)
	fmt.Println("message =", message)
	fmt.Println("err =", err)
	fmt.Printf("%#v received from %#v, message: %#v.\n", m.Title, message.NSQDAddress, string(message.Body))
	return
}

func init() {
	var (
		topic   = "topic_demo"
		channel = "first"
		config  = nsq.NewConfig()
		err     error
	)
	config.LookupdPollInterval = 15 * time.Second
	consumer, err = nsq.NewConsumer(topic, channel, config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("nsq.NewConsumer(%#v, %#v, %#v) return: (%#v, %#v).\n", topic, channel, config, consumer, err)
}

func main() {
	handler := &MyHandler{Title: "Yuan"}
	consumer.AddHandler(handler)
	addr := "127.0.0.1:4161"
	err := consumer.ConnectToNSQLookupd(addr)
	if err != nil {
		panic(err)
	}

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT)
	<-c
}
