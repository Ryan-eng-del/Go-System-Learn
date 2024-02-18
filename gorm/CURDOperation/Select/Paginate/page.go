package page

import (
	. "gorm/Model"

	"gorm.io/gorm"
)

func Paginate(page, pageSize int) func (db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		return db.Offset(page).Limit(pageSize)
	}
}

func Page() {
	c := []Article{}
	DB.Scopes(Paginate(1, 19)).Find(&c)
}