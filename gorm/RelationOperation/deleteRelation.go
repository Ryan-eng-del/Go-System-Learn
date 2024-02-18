package relation

import (
	"fmt"
	. "gorm/Model"
	"log"
)

func DeleteRelation () {
	// 添加一对多
	var a1 Author
	a1.Name = "max"

	if err := DB.Create(&a1).Error; err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("max id is ", a1.ID)

	var e1 Essay
	var e2 Essay

	e1.Subject = "e1"
	e2.Subject = "e2"

	if err := DB.Create([]*Essay{&e1, &e2}).Error; err != nil {
		log.Fatal(err)
	}

	fmt.Println("e1 id is ", e1.ID)
	fmt.Println("e2 id is ", e2.ID)

	// delete clear 一对多会将不包括在 replace 列表里的记录只为 NULL
	DB.Model(&a1).Association("Essays").Append([]*Essay{&e1, &e2})
	DB.Model(&a1).Association("Essays").Delete([]*Essay{&e1})
	DB.Model(&a1).Association("Essays").Clear()
	fmt.Println(len(a1.Essays))


		// 添加多对多
		var t1 Tag
		var t2 Tag
		var t3 Tag
	
		t1.Color = "white"
		t2.Color = "black"
		t3.Color = "dark"
	
		DB.Create([]*Tag{&t1, &t2, &t3})
	
		fmt.Println("t1 id is ",t1.ID)
		fmt.Println("t2 id is ", t2.ID)
		fmt.Println("t3 id is ",t3.ID)
	
	  // delete, clear 多对多会将不包括在 delete 列表里的记录清空
		DB.Model(&e1).Association("Tags").Append([]*Tag{&t1,&t3})
		// DB.Model(&e1).Association("Tags").Delete([]*Tag{&t1})
		DB.Model(&e1).Association("Tags").Clear()

		var am EssayMate
		var am1 EssayMate
	
		am.Title = "one meta"
		am1.Title = "one meta"
		DB.Create(&am)
		DB.Create(&am1)
	
		fmt.Println("essay meta id is ", am.ID)
		fmt.Println("essay meta1 id is ", am1.ID)
	
		// delete, clear 一对一 会将不包括在 delete 列表里的记录置为 NULL
		DB.Model(&e1).Association("EssayMate").Append(&am)
		DB.Model(&e1).Association("EssayMate").Delete(&am)
		DB.Model(&a1).Association("EssayMate").Clear()
		fmt.Println("meta relate essay id is ", *am.EssayID)

}