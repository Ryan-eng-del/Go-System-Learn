package list

import (
	"context"
	. "goRedis/redis"
)
func Index() {
	ctx := context.Background()
	Client.LIndex(ctx, "subjects", 1)
	Client.LSet(ctx, "subjects", 0, "0-value")
	Client.LInsert(ctx, "subjects", "BEFORE", "0-value", "-1-value")
	Client.LInsertBefore(ctx, "subjects", "0-value", "-1-value")
	// 获取匹配元素的第一个索引
	// Client.LPos(ctx, "subjects", "0-value")

	// 给予索引裁剪元素, 留下 【0-2】的元素
	Client.LTrim(ctx, "subjects",0, 2)
	// 0 删除全部
	// n >0 从头函数删除 n 个
	// -n <0 从尾部删除 n 个
	Client.LRem(ctx, "subjects", 1, "0-value")
}