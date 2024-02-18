package selectSend

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/IBM/sarama"
)


func SelectSend() {
	broker := []string{"localhost:9092"}

	conf := sarama.NewConfig()
	conf.Producer.Return.Successes = true

	producer, err := sarama.NewAsyncProducer(broker, conf)

	if err != nil {
		log.Fatal(err)
	}

	defer func () {
		if err := producer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	successCount, sendCount, errorCount := 0, 0, 0
	loop:
		for {
			time.Sleep(1 * time.Second)
			select {
			case producer.Input() <- &sarama.ProducerMessage{
				Topic: "async_topic",
				Value: sarama.StringEncoder("Hello, world!"),
			}:
				sendCount++
			case err := <- producer.Errors():
				log.Println(err)
				errorCount++
			case success := <- producer.Successes():
				fmt.Println("success", success)
				successCount++
			case <- signals:
				break loop
			}
		}

		fmt.Printf("successCount: %d, sendCount: %d, errorCount: %d", successCount, sendCount, errorCount)
}