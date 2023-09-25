package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/rs/zerolog"
)

/*
	Cache handler has a redis connection and store the short url as the redis key and long url as its value
	For example: we has short url:7lV3DJbBH and long url:https://p30download.com,
	so we must store the key with 7lV3DJbBH value and store its value https://p30download.com
*/

type CacheHandler interface {
	Get(ctx context.Context, shortUrl string) (string, error)
	Add(ctx context.Context, shortUrl, longUrl string) error
	Delete(ctx context.Context, shortUrl string) (bool, error)
}

type cacheHandler struct {
	client     *redis.Client
	log        *zerolog.Logger
	expireTime int
}

func NewCacheHandler(host, port string, expireTime int, log *zerolog.Logger) CacheHandler {
	redis_addr := fmt.Sprintf("%s:%s", host, port)
	fmt.Printf("redis_addr:%s\n", redis_addr)
	client := redis.NewClient(&redis.Options{
		Addr:     redis_addr,
		Password: "",
	})
	return &cacheHandler{
		client:     client,
		log:        log,
		expireTime: expireTime,
	}
}

func (c *cacheHandler) Get(ctx context.Context, shortUrl string) (string, error) {
	// check the `shortUrl` key from redis
	short, err := c.client.Get(ctx, shortUrl).Result()

	// check whether it's found or not
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		c.log.Error().Msgf("error in getting short url from redis: %s", err.Error())
		return "", err
	}

	return short, nil
}

func (c *cacheHandler) Add(ctx context.Context, shortUrl, longUrl string) error {
	// set the `shortUrl` as redis key and `longUrl` as its value and set the expiretime in seconds
	err := c.client.Set(ctx, shortUrl, longUrl, time.Duration(c.expireTime)*time.Second).Err()
	if err != nil {
		c.log.Error().Msgf("error in adding short and long urls to redis: %s", err.Error())
		return err
	}

	return nil
}

func (c *cacheHandler) Delete(ctx context.Context, shortUrl string) (bool, error) {
	// delete the cache key with the given `shortUrl`
	err := c.client.Del(ctx, shortUrl).Err()

	if err != nil {
		c.log.Error().Msgf("error in deleting key from redis: %s", err.Error())
		return false, err
	}

	return true, nil
}
