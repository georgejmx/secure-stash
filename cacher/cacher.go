package cacher

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Cacher struct {
	ctx context.Context
	rcache *redis.Client
}

func (c *Cacher) InitCacher() {
	c.ctx = context.Background()
	c.rcache = redis.NewClient(&redis.Options{
		Addr:	  "localhost:6379",
		Password: "",
		DB:		  0,  // use default DB
	})
}

func (c *Cacher) InsertKey(key, val string) error {
	err := c.rcache.Set(c.ctx, key, val, 0).Err()
	return err
}

func (c *Cacher) RetrieveKey(key string) (string, error) {
	val, err := c.rcache.Get(c.ctx, key).Result()
	return val, err
}