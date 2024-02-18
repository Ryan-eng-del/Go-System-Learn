package orm

import (
	"fmt"
	. "gorm/Model"
)
func GetSelect() {
	c := Article{}
	DB.Select("subject", "likes", "concat(subject, '-', likes) as V").First(&c, "id > ?", 1)
	fmt.Printf("%+v", c)

	c1 := []Article{}
	result := DB.Distinct("*").Find(&c1, "id > ?", 1)
	fmt.Printf("%+v", c1)
	fmt.Println(result.RowsAffected)
}
