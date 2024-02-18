package sortedset

import (
	"context"
	"fmt"
	. "goRedis/redis"
	"time"

	"github.com/go-redis/redis/v8"
)

func Del() {
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


	Client.ZRem(ctx, "iphone", "vivo", "iphone")
	Client.ZRemRangeByScore(ctx, "iphone", "2", "10")
	Client.ZRemRangeByRank(ctx, "iphone", 0, -1)

	Client.ZPopMin(ctx, "iphone", 1)
	Client.ZPopMax(ctx, "iphone", 1)
	
	// 支持阻塞弹出
	Client.BZPopMax(ctx, 1 * time.Second, "iphone")
	Client.BZPopMin(ctx,1 * time.Second, "iphone")

	fmt.Println(Client.ZRangeWithScores(ctx, "phones", 0, -1).Result())

}