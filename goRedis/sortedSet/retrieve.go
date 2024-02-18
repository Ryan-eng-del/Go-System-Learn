package sortedset

import (
	"context"
	. "goRedis/redis"
)

func Retrieve() {
	ctx := context.Background()
	// 获取成员分数
	Client.ZScore(ctx, "phones", "vivo")
	// 获取多个分值
	Client.ZMScore(ctx, "phones", "vio", "huawei")
	// false true 表示是否带有分值
	Client.ZRandMember(ctx, "phones", 2, false)
}