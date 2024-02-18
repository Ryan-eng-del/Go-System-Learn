package sortedset

import (
	"context"
	"fmt"
	. "goRedis/redis"

	"github.com/go-redis/redis/v8"
)

func Op() {
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

	// 交集 并集都是同理, 结果返回
	Client.ZInterWithScores(ctx, &redis.ZStore{
		Keys: []string{"phones", "things"},
		// 聚合的权重
		Weights: []float64{1, 1},
		// 聚合方式，默认是交集相加，可选值是 MAX MIN
		Aggregate: "SUM",
	})
	// 交集并存储到新的 key 中
	Client.ZInterStore(ctx, "new_phones", &redis.ZStore{
		
	})

	Client.ZUnionWithScores(ctx, redis.ZStore{
		Keys: []string{"phones", "things"},
		// 聚合的权重
		Weights: []float64{1, 1},
		// 聚合方式，默认是交集相加，可选值是 MAX MIN
		Aggregate: "SUM",
	})

	Client.ZDiffWithScores(ctx, "phones", "things")
	
	fmt.Println(Client.ZRangeWithScores(ctx, "phones", 0, -1).Result())

}