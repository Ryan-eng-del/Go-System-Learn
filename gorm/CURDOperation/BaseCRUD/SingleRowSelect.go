package orm

import (
	"fmt"
	. "gorm/Model"
	"reflect"
)

func SingleRowSelect() {
	
	firstArticle := Article{}
	DB.First(&firstArticle, "id > ?", 1)
	LastArticle := Article{}
	DB.Last(&LastArticle, "id > ?", 1)
	TakeArticle := Article{}
	DB.Take(&TakeArticle, "id > ?", 1)
	m1 := map[string]interface{}{}
	// map 类型有好处是，对于 select 到的一些额外字段，直接可以获取
	// 不需要在模型当中加入其余字段 concat("filed1", "filed2") as v1,
 	DB.Model(&Article{}).First(&m1, "id > ?", 1)
	FindOneArticle := Article{}
	DB.Limit(1).Find(&FindOneArticle, "id > ?", 1)
	FindArticle := Article{}
	DB.Find(&FindArticle, "id > ?", 1)
	fmt.Println(firstArticle.ID)
	fmt.Println(LastArticle.ID)
	fmt.Println(TakeArticle.ID)
	fmt.Println(FindOneArticle.ID)
	fmt.Println(FindArticle.ID)
	
	m := map[string]any{
		"str": 1,
	}
	i, ok := m["str"].(int)
	if ok {
		fmt.Println(reflect.TypeOf(i))
	}
}