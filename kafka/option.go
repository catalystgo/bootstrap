package kafka

import (
	"time"

	"github.com/IBM/sarama"
)

// Sarama Options

type SaramaOption func(*sarama.Config)

// Consumer Group Options

type ConsumerGroupOption func(*sarama.Config)

func WithRebalanceStrategy(strategy sarama.BalanceStrategy) ConsumerGroupOption {
	return func(c *sarama.Config) {
		c.Consumer.Group.Rebalance.Strategy = strategy
	}
}

func WithNewestOffset() ConsumerGroupOption {
	return func(c *sarama.Config) {
		c.Consumer.Offsets.Initial = sarama.OffsetNewest
	}
}

func WithOldestOffset() ConsumerGroupOption {
	return func(c *sarama.Config) {
		c.Consumer.Offsets.Initial = sarama.OffsetOldest
	}
}

// Sync Producer Options

type SyncProducerOption func(*sarama.Config)

func WithRequiredAcks(acks sarama.RequiredAcks) SyncProducerOption {
	return func(c *sarama.Config) {
		c.Producer.RequiredAcks = acks
	}
}

// Async Producer Options

type AsyncProducerOption func(*sarama.Config)

// Message Options

type MessageOption func(*sarama.ProducerMessage)

// WithKey sets the key of the message.
// The key is used to determine the partition of the message using hash partitioning.
func WithKey(key []byte) MessageOption {
	return func(p *sarama.ProducerMessage) {
		p.Key = sarama.ByteEncoder(key)
	}
}

// WithPartition sets the partition of the message.
func WithPartition(partition int32) MessageOption {
	return func(p *sarama.ProducerMessage) {
		p.Partition = partition
	}
}

// WithHeader append a header to the message.
func WithHeader(key []byte, value []byte) MessageOption {
	return func(p *sarama.ProducerMessage) {
		p.Headers = append(p.Headers, sarama.RecordHeader{Key: key, Value: value})
	}
}

// WithHeaders sets the headers of the message.
func WithHeaders(headers []sarama.RecordHeader) MessageOption {
	return func(p *sarama.ProducerMessage) {
		p.Headers = headers
	}
}

// WithTimestamp sets the timestamp of the message. The timestamp is the number of milliseconds since the Unix epoch.
func WithTimestamp(timestamp int64) MessageOption {
	return func(p *sarama.ProducerMessage) {
		p.Timestamp = time.Unix(timestamp, 0)
	}
}
