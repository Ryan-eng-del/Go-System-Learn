package encoder

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/IBM/sarama"
)


type Event struct {
	Message string `json:"message"`
	ID  int `json:"id"`
	Time time.Time `json:"time"`
	Type string `json:"type"`
}



func Encoder () {
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
				ev := Event{}
				if err := json.Unmarshal(msg.Value, &ev); err != nil {
					log.Printf("consumer id is %d, time is %s, value is %s, type is %s", ev.ID, ev.Time, ev.Message, ev.Type)
				}
			case <- sig:
				break loop
			}
		}
}