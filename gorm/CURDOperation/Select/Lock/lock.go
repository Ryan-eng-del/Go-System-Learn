package lock

import (
	. "gorm/Model"

	"gorm.io/gorm/clause"
)
func Lock() {
	c := []Article{}
	// for update
	DB.Model(&Article{}).Clauses(clause.Locking{
		Strength: "UPDATE",
	}).Find(&c)
	// for share
	DB.Model(&Article{}).Clauses(clause.Locking{
		Strength: "SHARE",
	}).Find(&c)

}