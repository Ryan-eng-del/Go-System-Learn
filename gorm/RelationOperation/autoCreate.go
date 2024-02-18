package relation

import (
	. "gorm/Model"
)
func AutoCreate() {
	var t Tag
	t.Color = "blue"
	DB.Create(&t)
	// 声明式的声明模型结构时,创建父模型的同时,也会创建各个关联的子模型 
	a := Essay{
		Subject: "一个组合的 Save",
		Author: Author{Name: "ccc"},
		Tags: []Tag{
			{Color: "11"},
			t,
		},
	}	
	DB.Create(&a)
}