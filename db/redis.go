package db

import (
	"strconv"

	"github.com/go-redis/redis"
)

type RedisDb struct {
	client *redis.Client
}

func NewRedisDb() Db {
	return &RedisDb{
		client: redis.NewClient(&redis.Options{
			Addr:     "redis:6379",
			Password: "",
			DB:       0,
		}),
	}
}

func (db *RedisDb) Get(key string) (int, error) {
	str, err := db.client.Get(key).Result()

	if err != nil {
		return 0, err
	}

	val, err1 := strconv.Atoi(str)
	return val, err1
}

func (db *RedisDb) Set(key string, value int) error {
	return db.client.Set(key, value, 0).Err()
}

func (db *RedisDb) Name() string {
	return "redis"
}
