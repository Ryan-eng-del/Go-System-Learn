package relation

import (
	"fmt"
	. "gorm/Model"
	"log"
)

// Find 是找到字表之后继续通过 Find 找关联字表
func FindRelations () {
		var a1 Author
		a1.Name = "max"
	
		if err := DB.Create(&a1).Error; err != nil {
			log.Fatal(err)
		}
		
		fmt.Println("max id is ", a1.ID)
	
		var e1 Essay
	
		e1.Subject = "第一篇文章"
	
		if err := DB.Create([]*Essay{&e1}).Error; err != nil {
			log.Fatal(err)
		}
	
		fmt.Println("e1 id is ", e1.ID)
		
		var es []Essay
		// where 语句是对 Association 中的字段, 进行筛选的
		DB.Model(&a1).Where("title like ?", "%ti").Association("Essays").Replace([]*Essay{&e1})
		// replace 一对多会将不包括在 replace 列表里的记录只为 NULL
		DB.Model(&a1).Association("Essays").Find(&es)
		fmt.Printf("%+v %s", a1.Essays, "查询前")
		DB.Model(&a1).Association("Essays").Find(&a1.Essays)
		fmt.Printf("%+v", a1.Essays)
		fmt.Printf("%+v", es)

		// 查询关联count
		count  := DB.Model(&a1).Association("Essays").Count()
		fmt.Println("作者关联的文章数目是", count)


		// 普通 find, 不会去查任何关联的模型
		var a2 Author
		DB.First(&a2, a1.ID)
		var e3 Essay
		DB.First(&e3, e1.ID)
		fmt.Println("单独使用 id 去查询 author")
		fmt.Printf("%+v", a2)
		fmt.Println("单独使用 id 去查询 essay")
		fmt.Printf("%+v", e3)
		fmt.Println(*e3.AuthorID)
}