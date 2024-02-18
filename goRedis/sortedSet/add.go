package sortedset

import (
	"context"
	"fmt"
	. "goRedis/redis"

	"github.com/go-redis/redis/v8"
)

func Add() {
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

	Client.ZAddArgs(ctx, "phones", redis.ZAddArgs{
		// 不存在添加，仅仅添加成员
		NX: false,
		// 存在添加，仅仅更新成员
		XX: false,
		// 只有当成员 Score < 原来的 Score， 才去更新
		LT: false,
		// 只有当成员 Score > 原来的 Score， 才去更新
		GT:  false,
		// 仅返回修改的成员数量
		Ch: false,
		Members: []redis.Z{
			{Score: 1, Member: "iphone"},
		},
	})


	fmt.Println(Client.ZRangeWithScores(ctx, "phones", 0, -1).Result())

}