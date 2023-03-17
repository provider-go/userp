package redis

import (
	"github.com/go-redis/redis"
	"github.com/provider-go/go-logger"
	"github.com/spf13/viper"
)

type REDIS struct {
	db *redis.Client // 数据库句柄
}

func Init() *REDIS {
	var db *REDIS
	var err error
	db, err = NewRedis(viper.GetString("redis.addr"), viper.GetString("redis.password"), viper.GetInt("redis.database"))
	if err != nil {
		logger.Error("Init", "step", "NewRedis", "err", err)
		return nil
	}
	return db
}

func NewRedis(addr, password string, db int) (*REDIS, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return &REDIS{db: client}, nil
}

// Set 写入数据
func (r *REDIS) Set(key, value string) {
	r.db.Set(key, value, 0)
}

// Get 读数据
func (r *REDIS) Get(key string) string {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("Redis Get", "step", "defer", "err", err)
		}
	}()
	value, err := r.db.Get(key).Result()
	if err != nil || value == "" {
		return ""
	}
	return value
}

// Del 删除数据
func (r *REDIS) Del(key string) {
	r.db.Del(key)
}
