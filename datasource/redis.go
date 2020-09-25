package datasource

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var Rds *redis.Client

func init() {
	//初始化redis 连接
	Rds = redis.NewClient(&redis.Options{
		Addr:"127.0.0.1:6379",
		Password:"",
		DB:0,
	})

	_, err := Rds.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}

func SaveToRedis(key string, data interface{}) (err error) {
	value, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("marshal data error:%v", err)
	}

	Rds.Set(context.Background(), key, value, 0)
	return
}

//func GetFromRedis(key string) (interface{}, error) {
//	value, err := Rds.Get(context.Background(), key).Result()
//	if err == redis.Nil {
//		return nil, fmt.Errorf("用户名不存在")
//	} else if err != nil {
//		panic(err)
//	}
//
//	var u handler.User
//	if err = json.Unmarshal([]byte(value), &u); err != nil {
//		return nil, err
//	}
//
//	return u, nil
//}