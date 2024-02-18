package str

import (
	"context"
	"fmt"
	. "goRedis/redis"
	"time"
)
func IncrDesc() {
	ctx := context.Background()
	Client.Set(ctx, "counter", "0", 10*time.Second)

	Client.Incr(ctx, "counter")
	Client.Incr(ctx, "counter")
	Client.IncrBy(ctx, "counter", 3)


	Client.Decr(ctx, "counter")
	Client.Decr(ctx, "counter")
	Client.DecrBy(ctx, "counter", 3)

	res := Client.Get(ctx, "counter")
	fmt.Println(res.Result())



}