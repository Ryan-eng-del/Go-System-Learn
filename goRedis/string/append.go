package str

import (
	"context"
	"fmt"
	. "goRedis/redis"
	"time"
)
func Append() {
	ctx := context.Background()
	Client.Set(ctx, "name", "WangBo", 10 * time.Second)
	Client.Append(ctx, "name", " is a boy")
	res := Client.Get(ctx, "name")
	fmt.Println(res.Result())

}