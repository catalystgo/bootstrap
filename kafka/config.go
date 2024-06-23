package kafka

import "github.com/IBM/sarama"

type ConfigOption func(*Config)

func WithClientID(clientID string) ConfigOption {
	return func(c *Config) {
		c.ClientID = clientID
	}
}

func WithVersion(version sarama.KafkaVersion) ConfigOption {
	return func(c *Config) {
		c.Version = version
	}
}

type Config struct {
	ClientID string
	Brokers  []string
	Version  sarama.KafkaVersion
}

func NewConfig(brokers []string, opts ...ConfigOption) *Config {
	c := &Config{
		ClientID: "CatalystGo", // default client ID
		Brokers:  brokers,
		Version:  sarama.V2_6_0_0, // default version
	}

	for _, o := range opts {
		o(c)
	}

	return c
}
