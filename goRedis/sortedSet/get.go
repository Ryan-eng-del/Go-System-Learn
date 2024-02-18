package sortedset

import (
	"context"
	"fmt"
	. "goRedis/redis"

	"github.com/go-redis/redis/v8"
)

func Get() {
	ctx := context.Background()
	Client.ZAdd(ctx, "phones", &redis.Z{
		Score: 100,
		Member: "iphone",
	}, &redis.Z{
		Score: 0,
		Member: "vivo",
	})

	Client.ZAdd(ctx, "phones", &redis.Z{
		Score: 10,
		Member: "vivo",
	})
	fmt.Println(Client.ZRange(ctx, "phones", 0, -1))
	fmt.Println(Client.ZRangeWithScores(ctx, "phones", 0, -1).Result())
	fmt.Println(Client.ZRangeArgsWithScores(ctx, redis.ZRangeArgs{
		Key: "iphone",
		// 默认为索引排序，start end 为索引区间 -1 表示从最后一个元素开始计数
		Start: nil, // (7  -inf 开区间和无穷的表示方法
		Stop: nil, // (7  -inf
		// 分数排序， start end 为分数区间 -inf +inf 表示正负无穷 ( 表示开区间
		ByScore: false,
		ByLex: false,
		Rev: false,
		Offset: 0,
		// 总计 0表示获取全部
		Count: 0,
	}))

	Client.ZCard(ctx, "phones")
	Client.ZCount(ctx, "phones", "10", "20")
	// 统计成员的索引位置
	Client.ZRank(ctx, "phones", "vivo")	
}