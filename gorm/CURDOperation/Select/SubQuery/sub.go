package sub

import (
	. "gorm/Model"
)

func Sub() {
  // from 型子查询
	subQuery := DB.Model(&Article{}).Select("id").Where("likes > ?", 10)
	var cs []Article
	DB.Where("author_id in (?)", subQuery).Find(&cs)

	var c []Article
	DB.Select("id").First(&c)
}