package update

import (
	"fmt"
	. "gorm/Model"
)

func Update() {
	var c Article

	// 主键更新
	// save 如果面对的是无主键的结构体,执行的是插入操作
	DB.Save(&c)
	// save 如果面对的是有主键的结构体,执行的是更新操作
	DB.Save(&c)

	// 条件更新
	m := map[string]interface{}{
		"likes": 1000,
		"author_id": 10,
	}

	result := DB.Model(&Article{}).Where("likes > ?", 100).Updates(m)

	// RowsAffected 是修改的记录数,不是 where 中满足条件的记录数, 因为有些不需要更新, 原本就是更新之后的值了
	fmt.Println(result.RowsAffected)
}