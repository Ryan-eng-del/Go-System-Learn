package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/IBM/sarama"
)


type Event struct {
	Message string `json:"message"`
	ID  int `json:"id"`
	Time time.Time `json:"time"`
	Type string `json:"type"`
}

type EventCoder Event

func (e EventCoder) Encode()([]byte, error) {
	val, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (e EventCoder) Length() int {
	val, err := json.Marshal(e)
	if err != nil {
		return 0
	}
	return len(val)
}

func GoRoutineSend() {
	broker := []string{"localhost:9092"}

	conf := sarama.NewConfig()
	conf.Producer.Return.Successes = true

	producer, err := sarama.NewAsyncProducer(broker, conf)

	successCount, sendCount, errorCount := 0, 0, 0

	if err != nil {
		log.Fatal(err)
	}


	wg := &sync.WaitGroup{}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	go func () {
		defer wg.Done()
		wg.Add(1)
		for suc := range producer.Successes() {
			log.Printf("partition is %d, offer is %d", suc.Partition, suc.Offset)
			successCount++
		}
	}()

	go func() {
		defer wg.Done()
		wg.Add(1)
		for suc := range producer.Errors() {
			fmt.Println(suc.Err)
			errorCount++
		}
	}()
	

	

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
  loop:
		for {
			time.Sleep(1 * time.Second)
			ev := Event{
				Message: "Hello world",
				Time: time.Now(),
				Type: "event",
				ID: 1,
			}
			select {
			case producer.Input() <- &sarama.ProducerMessage{
				// Topic: "go-routine",
				Topic: "more_partitions",
				Value: EventCoder(ev),
			}:
				sendCount++
			case <- sig:
				producer.Close()
				break loop
			}
		}
	wg.Wait()
	fmt.Printf("successCount: %d, sendCount: %d, errorCount: %d", successCount, sendCount, errorCount)
}
