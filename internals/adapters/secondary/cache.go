package adapters

import (
	"encoding/json"
	"log"

	"github.com/Erodot0/gym-memeber-management/internals/app/domains/ports"
	"github.com/go-redis/redis/v8"
)

type CacheServices struct {
	CacheClient *redis.Client
}

func (c *CacheServices) SetCache(data ports.CachePort) error {
	// Marshal the data
	bytes, err := json.Marshal(data)
	if err != nil {
		log.Printf("@SetCache: Error marshaling data: %v", err)
		return err
	}

	// Get the key and expiration
	key := data.SetCacheKey()
	expires := data.SetCacheExpiration()

	return c.CacheClient.Set(c.CacheClient.Context(), key, bytes, expires).Err()
}

func (c *CacheServices) GetCacheKeys(data ports.CachePort) ([]string, error) {
	// Get the key
	key := data.GetCacheKey()

	// Get the data
	return c.CacheClient.Keys(c.CacheClient.Context(), key).Result()
}

func (c *CacheServices) GetCacheFromKey(key string, data ports.CachePort) error {
	return c.CacheClient.Get(c.CacheClient.Context(), key).Scan(data)
}

func (c *CacheServices) GetCacheFromData(data ports.CachePort) error {
	// Get the key
	key := data.GetCacheKey()

	// Get the data
	return c.CacheClient.Get(c.CacheClient.Context(), key).Scan(data)
}

func (c *CacheServices) DelCache(data ports.CachePort) error {
	// Get the key
	key := data.GetCacheKey()

	// Delete the key
	return c.CacheClient.Del(c.CacheClient.Context(), key).Err()
}

func (c *CacheServices) DelCacheMultiple(data ports.CachePort) error {
	// Get the key
	key, err := c.GetCacheKeys(data)
	if err != nil {
		log.Printf("@DelCacheMultiple: Error getting keys: %v", err)
		return err
	}

	return c.CacheClient.Del(c.CacheClient.Context(), key...).Err()
}
