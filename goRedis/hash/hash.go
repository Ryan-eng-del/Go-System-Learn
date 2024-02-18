package hash

import (
	"context"
	"fmt"
	. "goRedis/redis"
)
func GetSetter() {
	ctx := context.Background()
	Client.HSet(ctx, "stu", "1", "wangbo", "2", "cyan", "3", "liyanyan")

	fmt.Println(Client.HGetAll(ctx, "stu").Result())

	fmt.Println(Client.HLen(ctx, "stu").Result())
	// count 为正数，不会重复抽取
	// count 为负数，会被重复抽取

	fmt.Println(Client.HRandField(ctx, "stu", 2, true).Result())
	fmt.Println(Client.HRandField(ctx, "stu", -2, true).Result())
	
	fmt.Println(Client.HExists(ctx, "stu", "1").Result())
	fmt.Println(Client.HExists(ctx, "stu", "5").Result())
	fmt.Println(Client.Do(ctx, "HSTRLEN", "stu", "1").Result())
}