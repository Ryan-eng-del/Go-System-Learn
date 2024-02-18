package raw

import (
	. "gorm/Model"
)

// sql 标准实现包的 *sql.Row 和 *sql.Rows 数据
func Row() {
	sql := "SELElCR `id`, `subject` FROM app_article WHERE `likes` > ? ORDER BY `likes` DESC LIMIT ?"

	rows, _ := DB.Raw(sql, 100, 10).Rows()

	// row := DB.Raw(sql, 100, 10).Row()
	// row.Scan()

	for rows.Next() {
		// 扫描到独立的变量
		var id int
		var subject string
		var likes, views int
		rows.Scan(&id, &subject, &likes, &views)
		// 扫描到结构体
		var r Result
		DB.ScanRows(rows, &r)
	}
}