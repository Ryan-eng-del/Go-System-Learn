package orm

import (
	"fmt"
	. "gorm/Model"
	"log"
	"time"
)

func CRUD() {
	article := &Article{
		// Subject: "Hello World",
		Likes: 0,
		Published: true,
		PublishTime: time.Now(),
		AuthorID: 42,
		BasicArticle: BasicArticle{
			AppCommon: 124,
		},
	}
	// article.AppCommon = 122

  // 创建
	if err := DB.Debug().Create(article).Error; err != nil {
		log.Fatalln(err)
	}

	// 获取
	article_retrieve := &Article{}
	article_retrieve.ID = article.ID
	if err := DB.First(&article_retrieve).Error; err != nil {
		log.Fatalln(err)
	}


	fmt.Println(article_retrieve.Likes, article_retrieve.UpdatedAt)
	article_retrieve.Likes = 100
	
	// 改
	if err := DB.Save(article_retrieve).Error; err != nil {
		log.Fatalln(err)
	}
	fmt.Println(article_retrieve.Likes, article_retrieve.UpdatedAt)

	// 删除
	if err := DB.Delete(article_retrieve).Error; err != nil {
		log.Fatal(err)
	}
}