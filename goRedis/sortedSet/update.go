package sortedset

import (
	"context"
	"fmt"
	. "goRedis/redis"

	"github.com/go-redis/redis/v8"
)

func Update() {
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

	// 10 代表加10， -10代表减10
	Client.ZIncrBy(ctx, "phones", -10, "vivo")
}