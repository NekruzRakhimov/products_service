package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/NekruzRakhimov/product_service/internal/config"
	"github.com/redis/go-redis/v9"
	"time"
)

type Cache struct {
	rdb *redis.Client
}

func NewCache(client *redis.Client) *Cache {
	return &Cache{
		rdb: client,
	}
}

func (c *Cache) Set(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	rawU, err := json.Marshal(value)
	if err != nil {
		return err
	}

	if err = c.rdb.Set(ctx, c.formatKey(key), rawU, duration).Err(); err != nil {
		fmt.Println("error during set:", err)
		return err
	}

	return nil
}

func (c *Cache) Get(ctx context.Context, key string, response interface{}) error {
	val, err := c.rdb.Get(ctx, c.formatKey(key)).Result()
	if err != nil {
		fmt.Println("error during get:", err)
		return err
	}

	if err = json.Unmarshal([]byte(val), response); err != nil {
		fmt.Println("error during unmarshal:", err)
		return err
	}

	return nil
}

func (c *Cache) formatKey(key string) string {
	return fmt.Sprintf("%s:%s", config.ServiceLabel, key)
}
