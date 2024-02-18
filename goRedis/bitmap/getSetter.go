package bitmap

import (
	"context"
	"fmt"
	. "goRedis/redis"

	"github.com/go-redis/redis/v8"
)
func GetterSetter() {
	ctx := context.Background()
	Client.SetBit(ctx, "userLog", 0, 1)
	Client.SetBit(ctx, "userLog", 1, 0)
	fmt.Println(Client.GetBit(ctx, "userLog", 0).Result()) 

	// get 获取未设置的位 为0
	// redis bitmap 是字符串值, 而 string value 的最大长度是 512MB， 由此可见，bitmap 可以存储 2*32 - 1位

	// 获取多少位是1
	// 以字节位单位，去统计，闭区间，默认是第一个字节
	fmt.Println(Client.BitCount(ctx, "userLog", &redis.BitCount{
		Start: 0,
		End: -1,
	}).Result())
}