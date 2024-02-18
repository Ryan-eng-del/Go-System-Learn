package str

import (
	"context"
	"fmt"
	. "goRedis/redis"
	"time"
)
func Sub() {
	ctx := context.Background()
	Client.Set(ctx, "name", "WangBo", 10 * time.Second)
	fmt.Println(Client.GetRange(ctx, "name", 0, -1).Result())
	fmt.Println(Client.GetRange(ctx, "name", 0, 5).Result())
	fmt.Println(Client.GetRange(ctx, "name", 0, 99).Result())
	fmt.Println(Client.GetRange(ctx, "name", -12, -1).Result())
	fmt.Println(Client.GetRange(ctx, "name", -12, -13).Err())
	fmt.Println(Client.SetRange(ctx, "name", 0, "Li").Result())
	fmt.Println(Client.Get(ctx, "name").Result())
}