package count

import . "gorm/Model"

func Count() {
	a := []Article{}
	var count int64
	DB.Find(&a).Count(&count)
}