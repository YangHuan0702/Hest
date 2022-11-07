package init

import (
	"github.com/garyburd/redigo/redis"
)

var RedisClient redis.Conn

func init() {
	rClient, err := redis.Dial("tcp", "localhost:6379", redis.DialPassword("123456"))
	if err != nil {
		panic(err)
	}
	RedisClient = rClient

	initGlobalId()
}

func initGlobalId() {
	id, err := redis.Int(RedisClient.Do("GET", "ID"))
	if err != nil {
		panic(err)
	}
	if id == 0 {
		RedisClient.Do("SET", "ID", 1000000)
	}
}
