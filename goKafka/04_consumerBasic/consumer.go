package consumerSinglePartition

import (
	"log"
	"os"
	"os/signal"

	"github.com/IBM/sarama"
)


func SingleConsumer () {
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Fatal(err)
	}
	
	defer func () {
		if err := consumer.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	
	partitionConsumer, err := consumer.ConsumePartition("go-routine", 0, sarama.OffsetNewest)

	if err != nil {
		log.Fatal(err)
	}

	defer func (){
		if err := partitionConsumer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	loop:
		for {
			select {
			case msg := <- partitionConsumer.Messages():
				log.Printf("consumer partition is %d, offset is %d, value is %s", msg.Partition, msg.Offset, string(msg.Value))
			case <- sig:
				break loop
			}
		}
}