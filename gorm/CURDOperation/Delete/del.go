package del

import . "gorm/Model"

func Dele() {

	// 条件删除
	DB.Delete(&Article{}, "likes < ?", 100)
	DB.Where("likes < ?", 100).Delete(&Article{})

	// 主键删除
	DB.Delete(&Article{}, 1)

	
	// 逻辑删除 | 硬删除
	cs := []Article{}
	DB.Unscoped().Find(&cs)
	DB.Unscoped().Delete(&Article{}, 1)
}