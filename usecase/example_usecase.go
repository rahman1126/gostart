package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"gostart/dto"
	"gostart/repository"
	"gostart/storage"
	"gostart/utils/cache"
	"gostart/utils/common"
	"gostart/utils/logger"
	"log"
	"time"
)

type ExampleUsecase struct {
	exampleRepository repository.ExampleRepositoryI
	ctxTimeout        time.Duration
	redisConn         redis.Conn
}

func NewExampleUsecase(er repository.ExampleRepositoryI, timeout time.Duration, redis redis.Conn) ExampleUsecaseI {
	return &ExampleUsecase{
		exampleRepository: er,
		ctxTimeout: timeout,
		redisConn: redis,
	}
}

func (eu ExampleUsecase) FindAll(ctx context.Context) ([]*dto.UserResponse, error)  {
	ctx, cancel := context.WithTimeout(ctx, eu.ctxTimeout)
	defer cancel()

	data := []*storage.User{}
	cached, err := cache.GET(eu.redisConn, fmt.Sprintf("%v", common.GET_ALL_USERS))
	if err != nil {
		logger.Info(nil, "usecase.FindAll: Get data from DB")
		data, err = eu.exampleRepository.FindAll(ctx)
		if err != nil {
			logger.Error(nil, "usecase.FindAll: Failed get data from DB")
			return nil, err
		}
		logger.Info(nil, "usecase.FindAll: Successfully get data from DB")
		_ = cache.SET(eu.redisConn, fmt.Sprintf("%v", common.GET_ALL_USERS), data)
	} else {
		err = json.Unmarshal([]byte(cached), &data)
		if err != nil {
			log.Println(err)
		}
	}

	if len(data) < 1 {
		logger.Warn(nil, "Cannot find any user")
	}

	res := make([]*dto.UserResponse, len(data))

	for i:=0; i<len(data); i++ {
		res[i] = &dto.UserResponse{
			Name: data[i].Name,
			Email: data[i].Email,
			Phone: data[i].Phone,
		}
	}

	return res, nil
}

func (eu ExampleUsecase) Find(ctx context.Context, id int) (*dto.UserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, eu.ctxTimeout)
	defer cancel()

	data := &storage.User{}
	cached, err := cache.GET(eu.redisConn, fmt.Sprintf("%v-%v", common.GET_A_USER, id))
	if err != nil {
		logger.Info(nil, "usecase.Find: Get data from DB")
		data, err = eu.exampleRepository.Find(ctx, id)
		if err != nil {
			logger.Error(nil, "usecase.Find: Failed get data from DB")
			return nil, err
		}
		logger.Info(nil, "usecase.Find: Successfully get data from DB")
		_ = cache.SET(eu.redisConn, fmt.Sprintf("%v-%v", common.GET_A_USER, id), data)
	} else {
		err = json.Unmarshal([]byte(cached), &data)
		if err != nil {
			log.Println(err)
		}
	}

	if data.ID == 0 {
		logger.Warn(nil, "usecase.Find: Cannot find user with id", id)
	}

	res := &dto.UserResponse{
		Name:      data.Name,
		Email:     data.Email,
		Phone:     data.Phone,
	}

	return res, nil
}