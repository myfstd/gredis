package gredis

import (
	"context"
	"github.com/myfstd/gredis/core"
	"time"
)

var redisClient *core.Client

func NewClient(addr string, pwd string, db int) *core.Client {
	ops := core.Options{
		Addr:            addr,
		Password:        pwd,
		DB:              db,
		ReadTimeout:     10,               // socket读取超时时间
		WriteTimeout:    10,               // socket写超时时间
		MaxRetries:      3,                //操作失败最大重试次数，默认不重试
		MinRetryBackoff: 10 * time.Second, // 最大重试时间间隔, 默认是 512ms; -1 表示关闭
	}
	redisClient = core.NewClient(&ops)
	//fmt.Println(redisClient)
	return redisClient
}

func SetVal(key string, val interface{}, exp ...time.Duration) error {
	var expiration time.Duration = 0
	if len(exp) == 0 {
		expiration = exp[0]
	}
	return redisClient.Set(context.Background(), key, val, expiration).Err()
}
func GetVal(key string) (interface{}, error) {
	return redisClient.Get(context.Background(), key).Result()
}
func Del(key string) error {
	return redisClient.Del(context.Background(), key).Err()
}

func HSetVal(key string, val interface{}, exp ...time.Duration) error {
	err := redisClient.HSet(context.Background(), key, val).Err()
	if len(exp) != 0 {
		redisClient.Expire(context.Background(), key, exp[0])
	}
	return err
}
func HGetVal(key string, field string) (interface{}, error) {
	return redisClient.HGet(context.Background(), key, field).Result()
}
func HGetAll(key string) (interface{}, error) {
	return redisClient.HGetAll(context.Background(), key).Result()
}
func HDel(key string, field string) {
	redisClient.HDel(context.Background(), key, field)
}

func LPushVal(key string, vals ...interface{}) error {
	return redisClient.LPush(context.Background(), key, vals).Err()
}
func LPullVal(key string) ([]string, error) {
	return redisClient.LRange(context.Background(), key, 0, -1).Result()
}
func LDel(key string, val interface{}) error {
	return redisClient.LRem(context.Background(), key, 0, val).Err()
}
