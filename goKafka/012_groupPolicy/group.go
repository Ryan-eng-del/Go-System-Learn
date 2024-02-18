package group

import "github.com/IBM/sarama"


func Group() {
	conf := sarama.NewConfig()
	// range 策略
	// 根据分区数量，确定每个消费者几个分区，多余的分区，前面的分区先分配

	// 2 topic 每个 topic 10 partitions 3 consumer
	// (2 * 10)/3 = 6 余下 2
	// 所以结果是
	// 第一个分区 6 + 2, 其他分区 6
	// 缺点是如果 partitions 数量多，那么有时第一个消费者分配的分区，远远多于其他消费者 


	// 轮询 不用说，一个一个轮着来

	// sticky 比较复杂的算法
	// 要达到的目的是 要尽可能的分配均匀，多个消费者主题分区数最大相差一个
	// 分区的分配要尽可能的和上次分配的保持相同，尽量保证原来分配的分区还在

	// 找到第一个，组内全部的 consumer 客户端都支持的策略，作为分配策略，因为每个客户端，支持的策略不一样
	conf.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{
		sarama.NewBalanceStrategySticky(),
		sarama.NewBalanceStrategyRange(),
		sarama.NewBalanceStrategyRoundRobin(),
	}

}