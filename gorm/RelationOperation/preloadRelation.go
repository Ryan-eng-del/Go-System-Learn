package relation

import (
	"fmt"
	. "gorm/Model"

	"gorm.io/gorm/clause"
)

// 多级 Preload, 链式 Preload
// Preload 不同于 Find 的是前者是查询父表, 同时加载字表, 后者是查到了父表, 再从父表查询字表
func PreloadOperation() {
	e1 := PreCreate()
	var e2 Essay

	// 条件也是针对于字表的条件
	DB.Preload("Author", "id IN ?", []uint{1, 2, 3}).Preload("EssayMate").Preload("Tags.Essays").First(e2, e1.ID)

	// 全部的关联都预加载, 只会加载第一层级的关联, 可以配合多级预加载一起使用
	DB.Preload(clause.Associations).First(&e2, 19)
}


func PreCreate() *Essay{
			// 添加一对多
			var a1 Author
			a1.Name = "max"
			DB.Create(&a1)
			var e1 Essay
			e1.Subject = "e1"
			DB.Create([]*Essay{&e1})
			
			DB.Model(&a1).Association("Essays").Append([]*Essay{&e1})
		
			// 添加一对一
			var am EssayMate
	
	
			am.Title = "one meta"
			DB.Create(&am)
		
			// 正向关联, 1对1没有反向关联, 会引起循坏引用
			DB.Model(&e1).Association("EssayMate").Append(&am)
		
			fmt.Printf("%v+", e1)
			// fmt.Printf("%v+", e2)
		
			// 添加多对多
			var t1 Tag
			var t2 Tag
			var t3 Tag
		
			t1.Color = "white"
			t2.Color = "black"
			t3.Color = "dark"
		
			DB.Create([]*Tag{&t1, &t2, &t3})
			// 正向关联
			DB.Model(&e1).Association("Tags").Append([]*Tag{&t1,&t3})
			return &e1
}