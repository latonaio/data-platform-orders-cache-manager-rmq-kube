package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

type Cache struct {
	rds *redis.Client
	log *logger.Logger
}

func NewCache(addr string, port interface{}, log *logger.Logger) *Cache {
	redisAddr := fmt.Sprintf("%s:%v", addr, port)
	opt := redis.Options{
		Addr: redisAddr,
		DB:   0,
		// Password: "",
	}

	cli := redis.NewClient(&opt)
	return &Cache{
		rds: cli,
		log: log,
	}
}

func (c *Cache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	res := c.rds.Set(ctx, key, value, expiration)
	if err := res.Err(); err != nil {
		return xerrors.Errorf("cache set error: %w", err)
	}
	c.log.Info("cache set to key %s", key)
	return nil
}

func (c *Cache) GetRaw(ctx context.Context, key string) ([]byte, error) {
	res := c.rds.Get(ctx, key)
	if err := res.Err(); err != nil {
		return []byte{}, xerrors.Errorf("cache get error: %w", err)
	}
	d, err := res.Bytes()
	if err != nil {
		return []byte{}, xerrors.Errorf("cache get bytes error: %w", err)
	}
	return d, nil
}
func (c *Cache) GetSlice(ctx context.Context, key string) ([]map[string]interface{}, error) {
	d := []map[string]interface{}{}
	b, err := c.GetRaw(ctx, key)
	if err != nil {
		return d, xerrors.Errorf("cache get raw error: %w", err)
	}
	err = json.Unmarshal(b, &d)
	if err != nil {
		return d, xerrors.Errorf("cache data unmarshal error: %w", err)
	}
	return d, nil
}
func (c *Cache) GetMap(ctx context.Context, key string) (map[string]interface{}, error) {
	d := map[string]interface{}{}
	b, err := c.GetRaw(ctx, key)
	if err != nil {
		return d, xerrors.Errorf("cache get raw error: %w", err)
	}
	err = json.Unmarshal(b, &d)
	if err != nil {
		return d, xerrors.Errorf("cache data unmarshal error: %w", err)
	}
	return d, nil
}
