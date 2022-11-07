package util

import (
	"Hest/init"
	"github.com/garyburd/redigo/redis"
)

func getId() int {
	id, err := redis.Int(init.RedisClient.Do("INTR", "ID"))
	if err != nil {
		panic(err)
	}
	return id
}

func getString(key string) string {
	v, err := redis.String(init.RedisClient.Do("GET", key))
	if err != nil {
		panic(err)
	}
	return v
}
