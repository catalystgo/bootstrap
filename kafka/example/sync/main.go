package main

import "github.com/catalystgo/bootstrap/kafka"

func main() {
	brokers := []string{"localhost:9094", "localhost:9095", "localhost:9096"}

	cfg := kafka.NewConfig(brokers, kafka.WithClientID("CatalystGo"))
	producer, err := kafka.NewSyncProducer(cfg)
	if err != nil {
		panic(err)
	}

	defer producer.Close()

	err = producer.Produce("test", []byte("Hello, World!"), kafka.WithKey([]byte("key")))
	if err != nil {
		panic(err)
	}

	println("SyncProducer Message sent", "✅")
}
