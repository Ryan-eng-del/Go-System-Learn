package orm

import (
	"database/sql"
	"fmt"
	. "gorm/Model"
)


func SingleColumnSelect() {
	var articles []sql.NullString
	DB.Model(&Article{}).Pluck("subject", &articles)

	for _, article := range articles {
		if article.Valid {
			fmt.Println(article.String)
		} else {
			fmt.Println("Null")
		}
	}

	var article_2 []string

	DB.Model(&Article{}).Pluck("concat(coalesce(subject, '[no subject]'), '-', likes)", &article_2)

	for _, a2 := range article_2 {
		fmt.Println(a2)
	}

}