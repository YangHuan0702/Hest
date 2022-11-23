package init

import (
	"github.com/garyburd/redigo/redis"
)

var RedisClient redis.Conn

func init() {
	// redis.DialPassword("123456")
	rClient, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialPassword("123456"))
	if err != nil {
		panic(err)
	}
	RedisClient = rClient

	initGlobalId()

	initDependencies()

}

func initDependencies() {

}

func initGlobalId() {
	id, err := RedisClient.Do("GET", "ID")
	if err != nil {
		panic(err)
	}
	i, err := redis.Int(id, err)
	if err != nil {
		return
	}
	if i == 0 {
		_, err := RedisClient.Do("SET", "ID", 1000000)
		if err != nil {
			panic("初始化全局ID异常")
		}
	}
}
