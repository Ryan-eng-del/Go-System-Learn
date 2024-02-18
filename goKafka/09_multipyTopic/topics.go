package topics

import (
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
)

func Topics() {
	broker := sarama.NewBroker("localhost:9092")

	conf := sarama.NewConfig()

	if err := broker.Open(conf); err != nil {
		log.Fatal(err)
	}
	// open 方法会连接，但是是非阻塞连接，不会等到连接成功了
	// 所以通常会强制连接判定
	connected, err := broker.Connected()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected", connected)

	topicKey := "more_partitions"

	topicDetail := &sarama.TopicDetail{
		NumPartitions: 3,
		ReplicationFactor: 1,
	}

	request := sarama.CreateTopicsRequest{
		TopicDetails: map[string]*sarama.TopicDetail{
			topicKey: topicDetail,
		},
		Timeout: 10 * time.Second,
		ValidateOnly: false,
	}


	res, err := broker.CreateTopics(&request)

	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.Version, "response")
}
