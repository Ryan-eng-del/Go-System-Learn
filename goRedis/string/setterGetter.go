package str

import (
	"context"
	"fmt"
	. "goRedis/redis"
	"time"

	"github.com/go-redis/redis/v8"
)

func StringSetter( ) {
	ctx := context.Background()

	// 设置
	// 有效期为 0, 数据会被持久化
	status := Client.Set(ctx, "name", "value", 2 * time.Second)
	status1 := Client.Set(ctx, "name1", "value1", 0)
	// 仅仅修改值, 不会重置有效期, 有效期还是按照原来的时间内执行
	status2 := Client.Set(ctx, "name", "value2", redis.KeepTTL)
	fmt.Println(status.Result())
	fmt.Println(status1.Result())
	fmt.Println(status2.Result())

	// 设置条件
	// NX key 不存在的时候设置
	// XX key  存在的时候设置
	Client.SetArgs(ctx, "name", "value3", redis.SetArgs{
		// "" 表示存在则更新, 不存在则添加, 是默认的模式
		Mode: "NX",
		// 有效期, 时间周期在2个小时内有效
		TTL: 0,
		// 是否返回原有值
		Get: false,
		// 有效期, 时间点
		ExpireAt: time.Time{},
		// 是否保持原有有效期
		KeepTTL: false,
	})

	// 同时设置多个
	Client.MSet(ctx, map[string]any{
		"map": "map1",
		"map2": "map2",
	})

	// 同时设置多个, 前提是全部的 key 都要不存在
	// Client.MSetNX()

	// 获取
	result := Client.Get(ctx, "name")
	val,  err := result.Result()
	// 如果key val已经过期,不存在,那么err 就为 redis:nil
	if err == redis.Nil {
		// key 不存在 字符串类型的 val 也是空字符串
		fmt.Println("key not exists")
	} else if val == "" {
		// key 存在, val 设置为空字符串
		fmt.Println("value is empty")
	}
	
	fmt.Println(val, err)
	r := Client.MGet(ctx, []string{"map", "map2"}...)
	v, _ := r.Result()
	fmt.Println("result", v)
	

	// 获取并删除 -> GetDel
	// 获取并设置有效期 -> GetEx
	// 获取多个 -> MGet
	// 获取字符串长度 -> StrLen
}	