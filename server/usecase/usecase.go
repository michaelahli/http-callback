package usecase

import (
	"encoding/json"
	"errors"
	"http-callback/helper"
	"http-callback/svcutil/cmd"
	"log"
	"time"

	"github.com/go-redis/redis"
)

type UC struct {
	Helper helper.Helper
	Bash   cmd.Terminal
	Redis  *redis.Client
}

// StoreToRedis save data to redis with key key
func (uc UC) StoreToRedis(key string, val interface{}) error {
	b, err := json.Marshal(val)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = uc.Redis.Set(key, string(b), 0).Err()

	return err
}

// StoreToRedisExp save data to redis with key and exp time
func (uc UC) StoreToRedisExp(key string, val interface{}, duration string) error {
	dur, err := time.ParseDuration(duration)
	if err != nil {
		return err
	}

	b, err := json.Marshal(val)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = uc.Redis.Set(key, string(b), dur).Err()

	return err
}

// GetFromRedis get value from redis by key
func (uc UC) GetFromRedis(key string, cb interface{}) error {
	res, err := uc.Redis.Get(key).Result()
	if err != nil {
		return err
	}

	if res == "" {
		return errors.New("[Redis] Value of " + key + " is empty.")
	}

	err = json.Unmarshal([]byte(res), &cb)
	if err != nil {
		return err
	}

	return err
}

// RemoveFromRedis remove a key from redis
func (uc UC) RemoveFromRedis(key string) error {
	return uc.Redis.Del(key).Err()
}
