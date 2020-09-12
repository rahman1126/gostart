package cache

import (
	"encoding/json"
	"errors"
	"github.com/gomodule/redigo/redis"
	"gostart/utils/conf"
	"gostart/utils/logger"
	"log"
)

func Pool() redis.Conn {
	if conf.IsUsingRedis() {
		pool := &redis.Pool{
			MaxIdle:   80,
			MaxActive: 12000,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", ":6379")
				if err != nil {
					//panic(err.Error())
					return nil, err
				}
				return c, err
			},
		}
		return pool.Get()
	}
	return nil
}

func Ping(c redis.Conn) error {
	if conf.IsUsingRedis() {
		_, err := c.Do("PING")
		if err != nil {
			logger.Error(nil,"cache.Ping: Unable to PING redis", err)
			return err
		}
		logger.Info(nil,"cache.Ping: Successfully PING redis")
		return nil
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