package where

import (
	"database/sql"
	. "gorm/Model"
)

func WhereType() {
	c := []Article{}
	condA := DB.Where("likes > ?", 10).Or("likes <= ?", 100)
	condB := DB.Where("likes > ? ", 10).Where(DB.Where("likes > ?", 10).Or("subject like ?", "gorm%"))
	query := DB.Where(condA).Where(condB)
	query.Find(&c)

	// 也可以使用 map 和 struct 来构建 where 查询语句，但还是推荐字符串，比较灵活，前者只能执行 And 逻辑
	c1 := []Article{}
	q := DB.Where("likes > @likes AND subject like @subject", sql.Named("subject", "sub%"), sql.Named("likes", 1))
	q.Find(&c1)


	c2 := []Article{}
	q1 := DB.Where("likes > @likes AND subject like @subject", map[string]interface{}{
		"likes": 1,
		"subject": "sub%",
	})
	q1.Find(&c2)
	
}