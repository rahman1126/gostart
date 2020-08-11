package cache

import (
	"encoding/json"
	"errors"
	"github.com/gomodule/redigo/redis"
	"gostart/utils/conf"
	"gostart/utils/logger"
	"log"
)

func Conn() redis.Conn {
	if conf.IsUsingRedis() {
		c, err := redis.Dial("tcp", conf.GetRedisAddr())
		if err != nil {
			logger.Error(nil,"cache.Conn: Connect to redis error", err)
			return nil
		}
		logger.Info(nil,"cache.Conn: Redis caching is on")
		return c
	}
	return nil
}

func SET(c redis.Conn, key string, value interface{}) (err error) {
	if conf.IsUsingRedis() {
		logger.Info(nil, "cache.SET: Set data to redis")
		byte, err := json.Marshal(value)
		if err != nil {
			log.Println(err)
			return err
		}
		_, err = c.Do("SET", key, byte)
		if err != nil {
			logger.Error(nil,"cache.SET: Failed set data to redis: ", err)
			return err
		}
		logger.Info(nil, "cache.SET: Successfully set data to redis")
		return nil
	} else {
		return errors.New("redis caching is OFF")
	}
}

func GET(c redis.Conn, key string) (value string, err error) {
	if conf.IsUsingRedis() {
		logger.Info(nil, "cache.GET: Get data from redis")
		value, err = redis.String(c.Do("GET", key))
		if err != nil {
			logger.Error(nil, "cache.GET: Failed get data from redis: ", err)
			return "", err
		}
		logger.Info(nil, "cache.GET: Successfully get data from redis")
		return value, nil
	} else {
		return value, errors.New("redis caching is OFF")
	}
}