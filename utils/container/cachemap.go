package container

import (
	"sync"
	"time"
)

/*
	带缓存的 map
*/

type (
	// CacheMap 带缓存的 Map
	CacheMap struct {
		sync.RWMutex
		cacheTime time.Duration
		Map       map[string]*val
	}
	val struct {
		t int64
		v interface{}
	}
)

// NewCacheMap ...
//	@param d 缓存时长
func NewCacheMap(d time.Duration) *CacheMap {
	c := &CacheMap{}
	c.cacheTime = d
	c.Map = make(map[string]*val)
	return c
}

// Set 设置值
//	@param k map 的 key
//	@param v map 的 value
func (c *CacheMap) Set(k string, v interface{}) {
	c.Lock()
	c.Map[k] = &val{time.Now().Add(c.cacheTime).Unix(), v}
	c.Unlock()
}

// Get 根据 key 获取 value
//	@return interface{} value 值
//	@return bool value 是否存在
func (c *CacheMap) Get(k string) (interface{}, bool) {
	now := time.Now().Unix()
	c.RLock()
	v, ok := c.Map[k]
	c.RUnlock()
	if !ok || now > v.t {
		return nil, false
	}
	return v.v, ok
}
