package handler

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/IBM/sarama"
)

func Handler() {
	addr := []string{"localhost:9092"}
	groupID := "group_1"
	conf := sarama.NewConfig()

	conf.Consumer.Return.Errors = true

	group, err := sarama.NewConsumerGroup(addr, groupID, conf)

	if err != nil {
		log.Fatal(err)
	}

	defer func () {
		group.Close()
	}()

	go func(){
		for err := range group.Errors() {
			log.Fatal(err)
		}
	}()

	topics := []string{"more_partitions"}

	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	go func() {
			defer wg.Done()
			wg.Add(1)
			for {
				fmt.Println("loop...")
				if err := group.Consume(ctx, topics, GroupConsumerHandler{}); err != nil {
					log.Fatal(err)
				}

				fmt.Println("code here")
				if ctx.Err() != nil {
					log.Println(ctx.Err())
					return
				}
			}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<- sig
	cancel()
	wg.Wait()

}

type GroupConsumerHandler struct {

}
// 组内消费者变化，会触发每个消费者（每个）会话的新的生命周期，会执行对应的 setup, comsumeClaim,cleanup

// setup 新会话开始时，可以通过 ConsumerGroupSession，来获得当前消费者的信息，这对当前消费者，最一些设置，例如起始消费的 offset
func (GroupConsumerHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

// 会话结束时，出发
func (GroupConsumerHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// 消费期间，一直阻塞触发
func (GroupConsumerHandler) ConsumeClaim(cgs sarama.ConsumerGroupSession, cgc sarama.ConsumerGroupClaim) error {
	for msg := range  cgc.Messages() {
		log.Printf("consumer partition is %d, offset is %d, value is %s", msg.Partition, msg.Offset, string(msg.Value))
		cgs.MarkMessage(msg, "")
	}
	return nil
}

