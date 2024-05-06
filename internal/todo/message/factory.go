package message

import (
	"encoding/json"
	"log/slog"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	kafkaconf "github.com/gsantosc18/todo/internal/todo/config/kafka"
)

func NewConsumer(kafkaConfig kafkaconf.KafkaProperties) (*kafka.Consumer, error) {
	topics := strings.Split(kafkaConfig.Topics, ",")
	consumer, consumerErr := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaConfig.Servers,
		"group.id":          kafkaConfig.GroupID,
		"auto.offset.reset": "earliest",
	})

	if consumerErr != nil {
		slog.Error("Has an error in consumer kafka", "error", consumerErr.Error())
		return nil, consumerErr
	}

	consumer.SubscribeTopics(topics, nil)

	return consumer, nil
}

func NewProducer(topic string, data any, kafkaconfig kafkaconf.KafkaProperties) error {
	producer, producerErr := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaconfig.Servers,
	})

	if producerErr != nil {
		slog.Error("Has an error on produce message to kafka", "error", producerErr.Error())
		return producerErr
	}

	message, serializeErr := json.Marshal(data)

	if serializeErr != nil {
		return serializeErr
	}

	producedErr := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          message,
	}, nil)

	return producedErr
}
