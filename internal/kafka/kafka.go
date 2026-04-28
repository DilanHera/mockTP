package kafka

import (
	"context"
	"crypto/tls"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

type KafkaAppConfig struct {
	Username string
	Password string
	Broker   string
	Topics   []string
}

type KafkaProducerConfig struct {
	Context  context.Context
	Messages []string
	Topic    string
}

type Kafka interface {
	Produce(config KafkaProducerConfig) error
	CloseWriters() error
}

type kafkas struct {
	kafkaBroker string
	Writers     map[string]*kafka.Writer
}

func NewKafka(config KafkaAppConfig) Kafka {
	config.Topics = []string{"sap.proxy.trackingStatusUpdated"}
	return &kafkas{
		kafkaBroker: config.Broker,
		Writers:     initWriters(config),
	}
}

func initWriters(config KafkaAppConfig) map[string]*kafka.Writer {
	writers := make(map[string]*kafka.Writer)
	for _, topic := range config.Topics {
		writers[topic] = &kafka.Writer{
			Addr:     kafka.TCP(config.Broker),
			Topic:    topic,
			Balancer: kafka.CRC32Balancer{},
			Transport: &kafka.Transport{
				TLS: &tls.Config{
					InsecureSkipVerify: true,
				},
				SASL: plain.Mechanism{Username: config.Username, Password: config.Password},
			},
		}
	}
	return writers
}

func (k *kafkas) CloseWriters() error {
	for _, w := range k.Writers {
		if err := w.Close(); err != nil {
			return err
		}
	}
	return nil
}
