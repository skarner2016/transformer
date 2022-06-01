package redis

import (
	"fmt"
	"transformer/library/config"
)
import "github.com/go-redis/redis/v8"

var (
	instantMap map[Instant]*redis.Client
)

func NewRedis(instant Instant) (*redis.Client, error) {
	if rdb, ok := instantMap[instant]; ok {
		return rdb, nil
	}

	rdbConf := new(RDBConfig)
	key := redisKeyPrefix + string(instant)
	if err := config.VipConfig.UnmarshalKey(key, &rdbConf); err != nil {
		return nil, err
	}

	addr := getAddr(rdbConf)

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: rdbConf.Pass,
		DB:       rdbConf.Database,
	})

	instantMap = make(map[Instant]*redis.Client)
	instantMap[instant] = rdb

	return rdb, nil
}

func getAddr(rdbConf *RDBConfig) string {
	return fmt.Sprintf("%s:%d", rdbConf.Host, rdbConf.Port)
}
