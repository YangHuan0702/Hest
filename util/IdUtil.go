package util

import (
	"github.com/garyburd/redigo/redis"
)

type OperatorUtil struct {
	RedisClient *redis.Conn
}

func (util *OperatorUtil) GetId() int {
	id, err := redis.Int((*util.RedisClient).Do("INTR", "ID"))
	if err != nil {
		panic(err)
	}
	return id
}

func (util *OperatorUtil) GetString(key string) string {
	v, err := redis.String((*util.RedisClient).Do("GET", key))
	if err != nil {
		panic(err)
	}
	return v
}
