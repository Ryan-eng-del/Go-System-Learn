package bitmap

import (
	"context"
	"fmt"
	. "goRedis/redis"
)
func AndOrNot() {
	ctx := context.Background()
	Client.SetBit(ctx, "keyOne", 0, 1)
	Client.SetBit(ctx, "keyTwo", 1, 1)

	// 与 或 非 异或
	Client.BitOpAnd(ctx, "deskKey", "keyOne", "keyTwo")
	Client.BitOpOr(ctx, "deskKey", "keyOne", "keyTwo")
	Client.BitOpNot(ctx, "deskKey", "keyOne")
	Client.BitOpXor(ctx, "deskKey", "keyOne", "keyTwo")

	// 获取 bitmap 结果
	fmt.Println(Client.GetBit(ctx, "deskKey", 0).Result())
	// 获取字符串结果，会转换成 acsii 字符
	fmt.Println(Client.Get(ctx, "deskKey"))
}