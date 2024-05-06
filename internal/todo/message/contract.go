package message

type ReceiverKafka interface {
	Topic() string
	Receiver(message []byte)
}
