package pkg

import (
	"fmt"
	"github.com/allegro/bigcache"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	DB *gorm.DB
	GlobalCache *bigcache.BigCache
)

func init()  {
	var err error
	DB, err = gorm.Open(mysql.Open("root:root@/casbin?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("database connnect falied, err: %v", err))
	}
	GlobalCache, err = bigcache.NewBigCache(bigcache.DefaultConfig(30 * time.Minute))
	if err !=  nil {
		panic(fmt.Sprintf("cache initialize failed, err : %v", err))
	}
}
