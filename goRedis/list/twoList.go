package list

import (
	"context"
	. "goRedis/redis"
)
func TwoList() {
	// 将一个队列的元素， push 到另外一个队列当中
	Client.RPopLPush(context.Background(), "list-src", "list-desc")
	Client.LMove(context.Background(), "list-src", "list-desc", "RIGHT", "LEFT")

}