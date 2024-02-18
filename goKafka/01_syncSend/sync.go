package sync

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)


func SyncSend(){
	brokers := []string{"localhost:9092"}


	producer, err := sarama.NewSyncProducer(brokers, nil)

	if err != nil {
		log.Fatal(err)
	}

	defer func (){
		if err := producer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	msg := &sarama.ProducerMessage{
		Topic: "sync_topic",
		Value: sarama.StringEncoder("Hello Kafka"),
	}

	partition, offset, err := producer.SendMessage(msg)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(partition, offset)
}