package db

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"io/ioutil"
)

func NewDBFromJSON(configFile string) *redis.Client {
	type redisConfig struct {
		Port     string `json:"port"`
		Password string `json:"password"`
		DB       int    `json:"DB"`
	}

	var rc redisConfig
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic("can not read redis config file")
	}
	json.Unmarshal(data, &rc)

	return redis.NewClient(&redis.Options{
		Addr:     rc.Port,
		Password: rc.Password,
		DB:       rc.DB,
	})
}