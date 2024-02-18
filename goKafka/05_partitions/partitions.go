package partitions

import (
	"fmt"
	"os"
	"os/signal"

	"log"

	"github.com/IBM/sarama"
)



func Partitions () {
		conf := sarama.NewConfig()
		consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, conf)

		if err != nil {
			log.Fatal(err)
		}

		defer func() {
			if err := consumer.Close(); err != nil {
				log.Fatal(err)
			}
		}()


		// partitions, err := consumer.Partitions("go-routine")
		partitions, err := consumer.Partitions("more_partitions")

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(partitions, "partitions")
	
		for _, partition := range partitions {
			// partitionConsumer, err := consumer.ConsumePartition("go-routine", partition, sarama.OffsetNewest)
			partitionConsumer, err := consumer.ConsumePartition("more_partitions", partition, sarama.OffsetNewest)

			if err != nil {
				log.Fatal(err)
			}

			go func(partitionConsumer sarama.PartitionConsumer){
				defer func ()  {
					if err := partitionConsumer.Close(); err != nil {
						log.Fatal(err)
					}
				}()

				for msg := range partitionConsumer.Messages() {
					log.Printf("consumer partition is %d, offset is %d, value is %s", msg.Partition, msg.Offset, string(msg.Value))
				}
				
			}(partitionConsumer)
		}
		sig := make(chan os.Signal, 1)

		signal.Notify(sig, os.Interrupt)
		<- sig
}