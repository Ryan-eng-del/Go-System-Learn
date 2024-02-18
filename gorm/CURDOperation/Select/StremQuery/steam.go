package stream

import (
	"fmt"
	. "gorm/Model"
	"log"
)
func Stream() {
	rows, err := DB.Model(&Article{}).Rows()

	if err != nil {
		log.Fatal(err)
	}

	defer func () {
		_ = rows.Close()
	}()

	for rows.Next() {
		c := Article{}
		DB.ScanRows(rows, &c)
		fmt.Println(c.ID)
	}
}