package orm

import (
	. "gorm/Model"

	"gorm.io/gorm"
)
func PrimaryKeySelect() {
	// 主键查询一条
	DB.First(&Article{}, 1)
	DB.Find(&Article{}, []int{1, 2})

	article := Article{
		Model: gorm.Model{
			ID: 1,
		},
	}
	// article.ID = 1
	DB.First(&article)
	// 主键字符串查询
	DB.First(&Article{}, "pk = ?", "stringPk")
	DB.Find(&Article{}, "pk IN ?", []string{"stringPK1", "stringPK2"})
}
