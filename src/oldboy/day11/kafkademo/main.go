package main

import (
	"fmt"
	"gopkg.in/Shopify/sarama.v1"
)

func main() {
	addr := []string{"127.0.0.1:9092"}
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	syncProducer, err := sarama.NewSyncProducer(addr, config)
	if err != nil {
		fmt.Printf("sarama.NewSyncProducer(%#v, %#v) return: (%#v, %#v).\n", addr, config, syncProducer, err)
	}
	defer func() { _ = syncProducer.Close() }()
}
