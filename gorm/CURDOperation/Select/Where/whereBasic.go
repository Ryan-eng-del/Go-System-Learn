package where

import (
	"fmt"
	. "gorm/Model"
)

func WhereMethod() {
	// And
	c := []Article{}
	r1 := DB.Find(&c, "likes > ? AND subject like ?", 0, "sub%")
	// fmt.Printf("%v", c)
	fmt.Println(r1.RowsAffected, c[0].ID)
	c1 := []Article{}
	query := DB.Where("likes > ?", 10)
	query.Where("subject LIKE ?", "sub%")
	r := query.Find(&c1)
	// fmt.Printf("%v", c1)
	fmt.Println(r.RowsAffected, c[0].ID)

	// Or
	c2 := []Article{}
	query = query.Or("id > ?", 5)
	query = query.Or(DB.Not("Id < ?", 8))
	query = query.Distinct("*")
	// query.Not 是对整个 query 语句进行否定 Not，而 DB.Not 是对括号当中的条件进行否定 
	r2 := query.Find(&c2)
	// fmt.Printf("%v", c1)
	fmt.Println(r2.RowsAffected, c2[0].ID)

	// Not
}