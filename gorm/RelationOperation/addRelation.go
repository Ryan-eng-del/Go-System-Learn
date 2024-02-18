package relation

import (
	"fmt"
	. "gorm/Model"
	"log"
)
func AddRelation () {
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
	
	// 指针可以获取到最新关联的父表id
	// 正向关联  正向关联 Author 表中的 Essays 是 Append 中的数据
	DB.Model(&a1).Association("Essays").Append([]*Essay{&e1, &e2, &e1})

	fmt.Println(len(a1.Essays))

	//? 如果此时这里新增加一个 author2 再去关联 e1 e2, 那么 e1, e2 的关联即被更新
	// 反向关联 反向关联 Author 表中的 Essays 是空数组
	// DB.Model(&e1).Association("Author").Append(&a1)
	// DB.Model(&e2).Association("Author").Append(&a1)

	fmt.Println("e1 relate author id is ", *e1.AuthorID)
	fmt.Println("e2 relate author id is ", *e2.AuthorID)

	fmt.Printf("%v+", a1)

	// 添加一对一
	var am EssayMate
	var am1 EssayMate

	am.Title = "one meta"
	am1.Title = "one meta"
	DB.Create(&am)
	DB.Create(&am1)

	fmt.Println("essay meta id is ", am.ID)
	fmt.Println("essay meta1 id is ", am1.ID)

	// 正向关联, 1对1没有反向关联, 会引起循坏引用
	DB.Model(&e1).Association("EssayMate").Append(&am)
	fmt.Println("meta relate essay id is ", *am.EssayID)
	// fmt.Println("meta relate essay id is ", *am1.EssayID)

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

	fmt.Println("t1 id is ",t1.ID)
	fmt.Println("t2 id is ", t2.ID)
	fmt.Println("t3 id is ",t3.ID)

	// 正向关联
	DB.Model(&e1).Association("Tags").Append([]*Tag{&t1,&t3})
	
	for _, t := range e1.Tags {
		fmt.Println("e1 relate tag id is",t.ID)
	}
	// 反向关联
	DB.Model(&t1).Association("Essays").Append(&e2)
	fmt.Println("e2 relate tag id is ", e2.Tags)
}