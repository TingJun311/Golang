package cache

import (
    "sync"
    "time"
)

type Cache struct {
    data *sync.Map
}

type CacheItem struct {
    value      interface{}
    expiration time.Time // when the time duration has been set it will auto expire and clear the cache
}

func NewCache() *Cache {
    return &Cache{
        data: &sync.Map{},
    }
}

func (c *Cache) Set(key string, value interface{}, expiration time.Duration) {
    item := CacheItem{
        value: value, 
        expiration: time.Now().Add(expiration),
    }
    c.data.Store(key, item)
}

func (c *Cache) Get(key string) (interface{}, bool) {
    if val, ok := c.data.Load(key); ok {
        item := val.(CacheItem)
        if time.Now().Before(item.expiration) {
            return item.value, true
        }
        c.data.Delete(key)
    }
    return nil, false
}

func (c *Cache) DisplayAll() map[string]interface{} {
    items := make(map[string]interface{})
    c.data.Range(func(key, value interface{}) bool {
        item := value.(CacheItem)
        if time.Now().Before(item.expiration) {
            items[key.(string)] = item.value
        } else {
            c.data.Delete(key)
        }
        return true
    })
    return items
}

// UpdateValue updates the value of a cache item with the specified key
func (c *Cache) Update(key string, value interface{}, expiration time.Duration) bool {
    if _, ok := c.data.Load(key); ok {
        newItem := CacheItem{
            value:      value,
            expiration: time.Now().Add(expiration),
        }
        c.data.Store(key, newItem)
        return true
    }
    return false
}


func (c *Cache) Delete(key string) {
    c.data.Delete(key)
}
