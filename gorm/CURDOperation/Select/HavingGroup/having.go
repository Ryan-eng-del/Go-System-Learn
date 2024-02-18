package having

import (
	"fmt"
	. "gorm/Model"
)

func Having() {

	type Result struct {
		TotalViews int
		AvgViews float64
		TotalLikes int
		AuthorId int
	}

	var rs []Result
	DB.Model(&Article{}).Select("author_id", "SUM(likes) as total_views", "AVG(likes) as avg_views").Group("author_id").Having("total_views > ? " ,10).Find(&rs)
	fmt.Printf("%+v", rs)

}