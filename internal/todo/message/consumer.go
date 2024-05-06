package message

import (
	"log/slog"

	"github.com/gsantosc18/todo/internal/todo/config/kafka"
)

var receivers []ReceiverKafka = []ReceiverKafka{}

func AddedReceiver(receiver ReceiverKafka) {
	receivers = append(receivers, receiver)
}

func StartConsumers() {
	configs := kafka.GetKafkaProperties()

	slog.Info("Connecting kafka")
	consumer, err := NewConsumer(*configs)

	if err != nil {
		slog.Error("Error when connecting kafka", "error", err.Error())
		return
	}

	defer consumer.Close()

	slog.Info("Connected to kafka server")

	consumers := make(map[string]ReceiverKafka)

	for r := range receivers {
		receiver := receivers[r]
		t := receiver.Topic()
		consumers[t] = receiver
	}

	for {
		message, consumerErr := consumer.ReadMessage(-1)

		if consumerErr != nil {
			slog.Error("Has an erro on consume kafka message", "error", consumerErr.Error())
			break
		}

		value := message.Value
		topic := message.TopicPartition.Topic

		receiver := consumers[*topic]
		receiver.Receiver(value)
	}
}
