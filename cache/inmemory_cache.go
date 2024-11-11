package cache

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
	"time"
)

type InMemoryCache struct {
	data map[string]cacheItem
	mu   sync.RWMutex
}

type cacheItem struct {
	value      interface{}
	expiration int64
}

func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{
		data: make(map[string]cacheItem),
	}
}

func (c *InMemoryCache) Set(key string, value interface{}, expiration time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = cacheItem{
		value:      value,
		expiration: time.Now().Add(expiration).Unix(),
	}
	return nil
}

func (c *InMemoryCache) Get(key string, dest interface{}) error {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, found := c.data[key]
	if !found || time.Now().Unix() > item.expiration {
		return fmt.Errorf("key not found")
	}

	// Verificar si `dest` es un puntero
	v := reflect.ValueOf(dest)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return errors.New("destination must be a non-nil pointer")
	}

	// Copiar el valor almacenado al valor apuntado por `dest`
	v.Elem().Set(reflect.ValueOf(item.value))
	return nil
}

func (c *InMemoryCache) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
	return nil
}
