package session

import (
	"fmt"
	. "gorm/Model"
	"log"

	"gorm.io/gorm"
)

func Session() {
	// session 解决同一个 DB 对象, 有子句重叠的问题  eg: where > 1 and where > 3
	// 对于以上的问题, 也可以让每一个查询从 DB 开始构建,会重新初始化 DB 对象, 但是相同的语句不能进行复用
	// 以上 session 出现了
	commonQuery := DB.Session(&gorm.Session{
		SkipHooks: true,
	}).Where("likes > ? ", 10)

	query1 := commonQuery.Where("id > ?", 1)
	fmt.Println("query1: ", query1)
	query2 := commonQuery.Where("id > ?", 10)
  fmt.Println("query2: ", query2)
	
	// dryRun 不真正执行sql, 可以查看 sql 语句
	db := DB.Session(&gorm.Session{
		DryRun: true,
	})
	
	stmt := db.Where("likes > ? ", 10).Statement
	log.Println(stmt.SQL.String())
	log.Println(stmt.Vars...)
}