package g_model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type BasicArticle struct {
	AppCommon int
}

type Article struct {
	gorm.Model
	BasicArticle `gorm:""`
	Subject      string `gorm:"type:varchar(255);not null;default:subject;comment:subject in article table"`
	Likes        uint   `gorm:"index"`
	Published    bool
	PublishTime  time.Time
	AuthorID     uint
	FirstName    string `gorm:"type:varchar(95);index:name,priority:2"`
	LastName     string `gorm:"type:varchar(100);index:name,priority:1"`
	FirstNameD   string `gorm:"type:varchar(95);index:name1"`
	LastNameD    string `gorm:"type:varchar(100);index:name1"`
	Age          uint8  `gorm:"index:age,sort:desc"`
	// 不需要迁移，并且只读，写被禁止
	V string `gorm:"-:migration;<-:false"`
}

// func (*Article) TableName() string {
// 	return "app_article"
// }

func (a *Article) AfterFind(db *gorm.DB) error {
	fmt.Println("a", a.ID)
	fmt.Println("Hook")
	return nil
}


type Author struct {
	gorm.Model
	Name string
	Email string
	Essays []Essay
}

type Essay struct {
	gorm.Model
	Subject string
	Likes uint
	Author Author
	// 一对多
	AuthorID *uint
	// 一对一
	EssayMate EssayMate
	// 多对多
	Tags []Tag `gorm:"many2many:content_tag"`
}

type EssayMate struct {
	gorm.Model	
	Title string
	EssayID *uint
}

type Tag struct {
	gorm.Model
	Color string
	Title string
	Essays []Essay `gorm:"many2many:content_tag"`
}


