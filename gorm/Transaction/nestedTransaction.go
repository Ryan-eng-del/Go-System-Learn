package transaction

import (
	. "gorm/Model"

	"gorm.io/gorm"
)

func NestedTransaction() {
	// 事务嵌套
	DB.Transaction(func(tx *gorm.DB) error {
		// 存储一个逻辑存储点
		tx.SavePoint("beforeA1")
		// 实现嵌套事务的复杂逻辑
		err := tx.Transaction(func(tx *gorm.DB) error {
			return nil
		})

		if err != nil {
			// 回滚到逻辑存储点
			tx.RollbackTo("beforeA1")
			tx.Transaction(func(tx *gorm.DB) error {
				return nil
			})
		}
		return nil
	})
}