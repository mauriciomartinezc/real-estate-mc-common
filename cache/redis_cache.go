package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(addr, password string, db int) *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	// Test the connection to Redis
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := client.Ping(ctx).Err(); err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}

	return &RedisCache{client: client}
}

// Set saves data in Redis cache with JSON serialization and a specified expiration
func (r *RedisCache) Set(key string, value interface{}, expiration time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Serialize the value to JSON
	jsonData, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("error marshaling data to JSON: %v", err)
	}

	// Set the data in Redis with expiration
	return r.client.Set(ctx, key, jsonData, expiration).Err()
}

// Get retrieves data from Redis cache and deserializes it into the provided destination
func (r *RedisCache) Get(key string, dest interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get the JSON data from Redis
	val, err := r.client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		// Key does not exist in Redis
		return fmt.Errorf("key not found")
	} else if err != nil {
		// Other Redis errors
		return fmt.Errorf("error retrieving data from Redis: %v", err)
	}

	// Deserialize JSON data into the destination
	if err := json.Unmarshal([]byte(val), dest); err != nil {
		return fmt.Errorf("error unmarshaling JSON data: %v", err)
	}
	return nil
}

// Delete removes a key from Redis cache
func (r *RedisCache) Delete(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Delete the key from Redis
	return r.client.Del(ctx, key).Err()
}
