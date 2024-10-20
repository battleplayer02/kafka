package main

import (
	"github.com/IBM/sarama"
)

func PushCommentToQueue(comment *Comment) {
	topic := "comments"
	brokerUrl := []string{"localhost:29092"}
	producer, err := ConnectProducer(brokerUrl)
}
func ConnectProducer(brokerUrl []string) sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	producer, err := sarama.NewSyncProducer(brokerUrl, config)
	if err != nil {
		panic(err)
	}
	return producer
}
