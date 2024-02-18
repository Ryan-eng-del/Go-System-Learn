package Ack

import (
	"log"

	"github.com/IBM/sarama"
)




func Ack() {
	conf := sarama.NewConfig()
	// 0 高吞吐量，对一致性的要求不高，不需要等任何分区，tcp ack 一返回就算成功
	conf.Producer.RequiredAcks = sarama.NoResponse
	// 1 AP 基本一致性+分区容忍，Leader 分区同步了，就算消息发送成功
	conf.Producer.RequiredAcks = sarama.WaitForLocal
	// -1 CP 强一致性+分区容忍 所有分区都同步了之后，才算消息发送成功
	conf.Producer.RequiredAcks = sarama.WaitForAll
	
	
	_ , err := sarama.NewAsyncProducer([]string{"localhost:9092"}, conf)

	if err != nil {
		log.Fatal(err)
	}

}