package kafka

import (
	"github.com/segmentio/kafka-go"
)

func (k *kafkas) Produce(config KafkaProducerConfig) error {
	w := k.Writers[config.Topic]

	messages := make([]kafka.Message, 0, len(config.Messages))
	for _, msg := range config.Messages {
		messages = append(messages, kafka.Message{Value: []byte(msg)})
	}
	return w.WriteMessages(config.Context, messages...)
}
