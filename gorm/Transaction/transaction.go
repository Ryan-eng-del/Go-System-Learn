package transaction

import (
	. "gorm/Model"
	"log"

	"gorm.io/gorm"
)
func Transaction() {
	// 事务开始 | 事务回滚 事务提交
	DB.Begin()
	DB.Rollback()
	DB.Commit()

	// gorm 维护事务操作, 其中只要 return error 事务就会回滚
	if err := DB.Transaction(func(tx *gorm.DB) error {
		if err := DB.Create(&Article{}).Error; err != nil {
			log.Fatal(err)
		}
		return nil
	}); err != nil {
		// 这里的错误只是用来提醒开发者, 事务失败了
		log.Println(err)
	}


	
}