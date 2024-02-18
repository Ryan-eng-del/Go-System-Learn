package ctx

import (
	"context"
	. "gorm/Model"
	"time"
)
func Ctx () {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	a := Article{}
	DB.WithContext(ctx).Where("likes > ?", 10).Find(&a)
}