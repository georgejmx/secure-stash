package cacher

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type CacherTemplate interface {
	InitCacher()
	RetrieveEntries() ([]string, error)
	InsertKey(key string, val []byte) error
	DetermineIfExists(key string) (bool, error)
	RetrieveEntry(key string) ([]byte, error)
}

type Cacher struct {
	ctx context.Context
	rcache *redis.Client
}

// Initialise program connection to redis
func (c *Cacher) InitCacher() {
	c.ctx = context.Background()
	c.rcache = redis.NewClient(&redis.Options{
		Addr:	  "localhost:6379",
		Password: "",
		DB:		  0,  // use default DB
	})
}

// Retrieves all entry keys from redis
func (c *Cacher) RetrieveEntries() ([]string, error) {
	cmd := c.rcache.Do(c.ctx, "KEYS", "*")
	return cmd.StringSlice()
}

// Add a new value to redis
func (c *Cacher) InsertKey(key string, val []byte) error {
	err := c.rcache.Set(c.ctx, key, val, 0).Err()
	return err
}

// Returns true if the specified key exists in Redis
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

// Retrieve an entry from redis
func (c *Cacher) RetrieveEntry(key string) ([]byte, error) {
	val, err := c.rcache.Get(c.ctx, key).Bytes()
	return val, err
}
