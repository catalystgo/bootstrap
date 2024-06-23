package main

import (
	"context"
	"time"

	"github.com/IBM/sarama"
	"github.com/catalystgo/bootstrap/kafka"
)

func main() {
	ctx := context.Background()
	brokers := []string{"localhost:9094", "localhost:9095", "localhost:9096"}

	cfg := kafka.NewConfig(brokers, kafka.WithClientID("CatalystGo"))
	consumerGroup, err := kafka.NewConsumerGroup(cfg, "cataystgo")
	if err != nil {
		panic(err)
	}

	go func() {
		syncProducer, err := kafka.NewSyncProducer(cfg)
		if err != nil {
			panic(err)
		}

		defer syncProducer.Close()

		time.Sleep(3 * time.Second)
		syncProducer.Produce("test", []byte("Hello, World!"))
	}()

	ch := make(chan struct{})
	err = consumerGroup.Consume(ctx, []string{"test"}, func(message *sarama.ConsumerMessage) error {
		println(string(message.Value))
		ch <- struct{}{}
		return nil
	})

	if err != nil {
		panic(err)
	}

	<-ch

	err = consumerGroup.Close()
	if err != nil {
		panic(err)
	}
}
