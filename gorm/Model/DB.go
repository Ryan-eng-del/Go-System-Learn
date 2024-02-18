package g_model

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)
var DB *gorm.DB

func init() {
  // refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
  dsn := "root:123456@tcp(127.0.0.1:3307)/wangbo?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "app_",
		},
	})

	if err != nil {
		log.Fatalln(err)
	}

	if err := db.AutoMigrate(&Article{}, &Author{}, &Tag{}, &Essay{}, &EssayMate{}); err != nil {
		log.Fatal(err)
	}
	fmt.Println("migrate success")
	DB = db
	
	DB.Logger = logger.Default.LogMode(logger.Info)
}
