package create

import (
	"fmt"
	. "gorm/Model"

	"gorm.io/gorm/clause"
)
func Conflict() {

	// InsertUpdate 全部更新
	c1 := Article{}
	c1.Subject = "111"
	// 默认主键冲突,会报错,设置为 updateAll 之后,对于冲突数据, 最新插入的 会更新旧的字段
	result1 := DB.Clauses(clause.OnConflict{UpdateAll: true}).Create(&c1)
	fmt.Println(result1.RowsAffected)


		// InsertUpdate 部分指定字段更新
		c2 := Article{}
		c2.Subject = "111"
		// 默认主键冲突,会报错,设置为 updateAll 之后,对于冲突数据, 最新插入的 会更新旧的字段
		result1 = DB.Clauses(clause.OnConflict{DoUpdates: clause.AssignmentColumns([]string{"likes", "subject"})}).Create(&c1)
		fmt.Println(result1.RowsAffected)
	
}