package cache

import (
	"sync"
	"time"
)

//type cache

type cahcheItem struct {
	value interface{}
	ttl time.Time
}

type cache struct {
	dt map[string]cahcheItem
	lcr sync.Mutex
}

func New()*cache{
	nCache := cache{dt:make(map[string]cahcheItem),lcr: sync.Mutex{}}
	go func() {
		for {
			nCache.lcr.Lock()
			for k, ch := range nCache.dt {
				if time.Now().After(ch.ttl) {
					delete(nCache.dt,k)
				}
			}
			nCache.lcr.Unlock()
		}
	}()
	return &nCache
}

func (c *cache) Set(key string,value interface{},timeLive time.Duration){
	c.lcr.Lock()
	defer c.lcr.Unlock()
	c.dt[key] = cahcheItem{
		value,
		time.Now().Local().Add(timeLive),
	}
}

func (c *cache) Get(key string)interface{}{
	c.lcr.Lock()
	defer c.lcr.Unlock()
	return c.dt[key].value
}

func (c *cache) Delete(key string){
	c.lcr.Lock()
	defer c.lcr.Unlock()
	delete(c.dt,key)
}