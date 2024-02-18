package create

import (
	"fmt"
	. "gorm/Model"
)
func Create() {

	// 单条插入
	c1 := Article{}
	c1.Subject = "111"
	result1 := DB.Create(&c1)
	fmt.Println(result1.RowsAffected)

	// 注意使用map插入的数据, create_at | update_time 不会更新
	m := map[string]any{
		"Subject": "122",
	}

	DB.Model(&Article{}).Create(m)

	// 多条插入 map 也是同理
	c2 := []Article{
		{Subject: "22"},
		{Subject: "11"},
	}
	result1 = DB.Create(&c2)
	fmt.Printf("result1.RowsAffected: %v\n", result1.RowsAffected)


	// 分批插入 两条两条的插入
	c3 := []Article{
		{Subject: "22"},
		{Subject: "11"},
	}
	result1 = DB.CreateInBatches(&c3, 2)
	fmt.Printf("result1.RowsAffected: %v\n", result1.RowsAffected)


}