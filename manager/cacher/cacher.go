package cacher

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type CacherTemplate interface {
	InitCacher()
	InsertKey(key string, val []byte) error
	DetermineIfExists(key string) (bool, error)
	RetrieveEntry(key string) ([]byte, error)
}

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

func (c *Cacher) InsertKey(key string, val []byte) error {
	err := c.rcache.Set(c.ctx, key, val, 0).Err()
	return err
}

/**
* Returns true if the specified key exists in Redis
*/
func (c *Cacher) DetermineIfExists(key string) (bool, error) {
	result, err := c.rcache.Exists(c.ctx, key).Result()
	if err != nil {
		return false, err
	}

	if result == 0 {
		return false, nil
	}

	return true, nil
}

func (c *Cacher) RetrieveEntry(key string) ([]byte, error) {
	val, err := c.rcache.Get(c.ctx, key).Bytes()
	return val, err
}