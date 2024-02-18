package raw

import (
	"fmt"
	. "gorm/Model"
)

type Result struct {
	ID int
	subject string
	likes int
}


func Raw() {
	sql := "SELElCR `id`, `subject` FROM app_article WHERE `likes` > ? ORDER BY `likes` DESC LIMIT ?"
	var rs []Result

	// Scan 是去将结果扫描道结果中, 比如select
	DB.Raw(sql, 100, 10).Scan(&rs)
	// Exec 只负责执行这条语句, 比如增加, 删除, 更新
	result := DB.Exec(sql, 100, 10)
	fmt.Println(result.RowsAffected)
}