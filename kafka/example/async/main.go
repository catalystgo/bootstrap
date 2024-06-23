package main

import "github.com/catalystgo/bootstrap/kafka"

func main() {
	brokers := []string{"localhost:9094", "localhost:9095", "localhost:9096"}

	cfg := kafka.NewConfig(brokers, kafka.WithClientID("CatalystGo"))
	producer, err := kafka.NewAsyncProducer(cfg)
	if err != nil {
		panic(err)
	}

	defer producer.Close()

	producer.Produce("test", []byte("Hello, World!"))
	println("AsyncProducer Message sent", "✅")
}
