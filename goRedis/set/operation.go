package set

import (
	"context"
	"fmt"
	. "goRedis/redis"
)
func Operation() {
	ctx := context.Background()
	Client.SAdd(ctx, "A", "1", "2", 1)
	Client.SAdd(ctx, "B", "1", "2")

	fmt.Println(Client.SMembers(ctx, "A").Result())

	// 交集，并集，补集
	Client.SInter(ctx, "A", "B")
	Client.SUnion(ctx, "A", "B")
	Client.SDiff(ctx, "A", "B")
} 

