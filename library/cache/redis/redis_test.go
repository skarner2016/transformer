package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"transformer/library/config"
	"testing"
)

func TestRedis_GetByInstant(t *testing.T) {
	config.InitConfig()

	rdb, err := NewRedis(InstantDefault)
	if err != nil {
		panic("get redis instant error")
	}

	rdb2, err := NewRedis(InstantDefault)
	if err != nil {
		panic("get redis instant error")
	}

	fmt.Println(fmt.Sprintf("%p", rdb))
	fmt.Println(fmt.Sprintf("%p", rdb2))
	fmt.Println([]*redis.Client{rdb, rdb2})
	fmt.Println(&rdb, &rdb2)

	fmt.Println(rdb.Get(context.Background(), "abc"))
}


