package init

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DbCli *gorm.DB
var RedisClient *redis.Conn

func init() {
	// 数据库
	initDB()

	// 初始化Redis
	initRedis()

	MvcInit(DbCli, RedisClient)
}

func initRedis() {
	// redis.DialPassword("123456")
	rClient, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialPassword("123456"))
	if err != nil {
		panic(err)
	}
	RedisClient = &rClient
	initGlobalId()
}

func initDB() {
	fmt.Printf("start DB Init ...\n")
	open, err := gorm.Open(mysql.Open("root:123456@tcp(localhost:3306)/Hest?parseTime=true&loc=Asia%2FShanghai"))
	if err != nil {
		panic(err)
	}
	db, _ := open.DB()
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(2)
	db.SetConnMaxIdleTime(time.Minute)
	DbCli = open
}

func initGlobalId() {
	id, err := (*RedisClient).Do("GET", "ID")
	if err != nil {
		panic(err)
	}
	i, err := redis.Int(id, err)
	if err != nil {
		return
	}
	if i == 0 {
		_, err := (*RedisClient).Do("SET", "ID", 1000000)
		if err != nil {
			panic("初始化全局ID异常")
		}
	}
}
