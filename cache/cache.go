package cache

import (
	"time"
)

type Cache interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string, dest interface{}) error
	Delete(key string) error
}
