package kafka

import (
	"context"

	"github.com/IBM/sarama"
	"github.com/catalystgo/logger/logger"
	"go.uber.org/zap"
)

type AsyncProducer interface {
	Produce(topic string, message []byte, opts ...MessageOption)
	ProduceWithContext(ctx context.Context, topic string, message []byte, opts ...MessageOption)
	Close() error
}

type asyncProducer struct {
	p sarama.AsyncProducer
}

func NewAsyncProducer(cfg *Config, opts ...AsyncProducerOption) (AsyncProducer, error) {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Version = cfg.Version
	saramaConfig.ClientID = cfg.ClientID
	saramaConfig.Producer.Return.Errors = true

	for _, opt := range opts {
		opt(saramaConfig)
	}

	producer, err := sarama.NewAsyncProducer(cfg.Brokers, saramaConfig)
	if err != nil {
		return nil, err
	}

	p := &asyncProducer{p: producer}

	go p.handleErrors()

	return p, nil
}

func (p *asyncProducer) Produce(topic string, message []byte, opts ...MessageOption) {
	p.ProduceWithContext(context.Background(), topic, message, opts...)
}

func (p *asyncProducer) ProduceWithContext(ctx context.Context, topic string, message []byte, opts ...MessageOption) {
	producerMessage := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(message),
	}

	for _, opt := range opts {
		opt(producerMessage)
	}

	select {
	case p.p.Input() <- producerMessage:
	case <-ctx.Done():
	}
}

func (p *asyncProducer) Close() error {
	return p.p.Close()
}

func (p *asyncProducer) handleErrors() {
	for err := range p.p.Errors() {
		logger.ErrorKV(context.Background(), "asyncProducer send message error",
			zap.Error(err.Err),
			zap.String("topic", err.Msg.Topic),
			zap.Int64("offset", err.Msg.Offset),
			zap.Int32("partition", err.Msg.Partition),
			zap.String("timestamp", err.Msg.Timestamp.String()),
		)
	}
}
