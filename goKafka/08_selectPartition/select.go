package partition

import "github.com/IBM/sarama"


func SelectPartition() {
	conf := sarama.NewConfig()

	// 生产者随机选择分区，均匀发送, 默认是随机的
	conf.Producer.Partitioner = sarama.NewRandomPartitioner
	// 生产者轮询选择分区
	conf.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	// 生产者hash选择分区，同一个 user Id 的标识，经过 hash 之后，在同一个分区
	conf.Producer.Partitioner = sarama.NewHashPartitioner
	// 手动指定选择分区
	conf.Producer.Partitioner = sarama.NewManualPartitioner
	// 自定义 Hash算法，来选择分区
	// conf.Producer.Partitioner = sarama.NewCustomHashPartitioner
	// 自定义选择分区
	// conf.Producer.Partitioner = sarama.NewCustomPartitioner()
}