package list

import (
	"context"
	"fmt"
	. "goRedis/redis"
	"time"
)
func PushPop() {
	ctx := context.Background()
	Client.LPushX(ctx, "subjects", "no exists")
	Client.LPush(ctx, "subjects", "Go", "Redis")
	Client.RPush(ctx, "subjects", "K8s")
	Client.LPop(ctx, "subjects")
	Client.BLPop(ctx, 10 *time.Second, "subjects")
	Client.LPopCount(ctx, "subjects", 1)

	fmt.Println(Client.LRange(ctx, "subjects",0, -1))




}