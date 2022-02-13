package kafka

import ckafka "github.com/confluentinc/confluent-kafka-go/kafka"

type Consumer struct {
	Topics    []string
	ConfigMap *ckafka.ConfigMap
}

func NewConsumer(configMap *ckafka.ConfigMap, topics []string) *Consumer {
	return &Consumer{
		ConfigMap: configMap,
		Topics:    topics,
	}
}

func (ref *Consumer) Consume(msgChan chan *ckafka.Message) error {
	consumer, err := ckafka.NewConsumer(ref.ConfigMap)
	if err != nil {
		panic(err)
	}

	err = consumer.SubscribeTopics(ref.Topics, nil)
	if err != nil {
		panic(err)
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			msgChan <- msg
		}
	}
}
