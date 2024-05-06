package kafka

import (
	"os"
)

type KafkaProperties struct {
	Servers         string
	GroupID         string
	AutoOffsetReset string
	Topics          string
}

func GetKafkaProperties() *KafkaProperties {
	kafkaProperties := KafkaProperties{
		Servers: os.Getenv("KAFKA_SERVERS"),
		GroupID: os.Getenv("KAFKA_GROUP"),
		Topics:  os.Getenv("KAFKA_TOPICS"),
	}
	return &kafkaProperties
}
