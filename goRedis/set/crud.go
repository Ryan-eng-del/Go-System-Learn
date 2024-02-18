package set

import (
	"context"
	. "goRedis/redis"
)
func CURD() {
	ctx := context.Background()
	// 增加
	Client.SAdd(ctx, "A", "1", "2", 1)
	Client.SAdd(ctx, "B", "1", "2")
	// 删除
	Client.SRem(ctx, "A", "2")
	Client.SPop(ctx, "A")
	// 获取·
	Client.SMembers(ctx,"À")
	Client.SRandMember(ctx, "A")
	Client.SRandMemberN(ctx, "A", 2)
	// 统计
	Client.SIsMember(ctx, "A", "a")
	Client.SCard(ctx, "A")

} 

