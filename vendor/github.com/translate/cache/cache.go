package cache

import (
	//"fmt"
	//"math/rand"
	"sync"
	"time"
)

type Cache struct {
	createAt time.Time
	life     time.Duration
	funcs    func() interface{}
	val      interface{}
}

func (t *Cache) isExpire() bool {
	if t.life == 0 {
		return false
	}

	return time.Now().Sub(t.createAt) > t.life
}

func (t *Cache) reflesh() error {
	t.val = t.funcs()
	t.createAt = time.Now()
	return nil
}

type CacheDict struct {
	sync.RWMutex
	CachePools map[string]*Cache
	interval   int
}

func NewCache() *CacheDict {
	CacheGvr := new(CacheDict)
	CacheGvr.CachePools = make(map[string]*Cache)
	CacheGvr.interval = 3

	//go CacheGvr.GC()

	return CacheGvr
}

func (d *CacheDict) Get(key string) interface{} {
	d.RLock()
	defer d.RUnlock()
	if value, ok := d.CachePools[key]; ok {
		if value.isExpire() {
			value.reflesh()
		}
		return value.val
	}
	return nil
}

func (d *CacheDict) Put(key string, funcs func() interface{}, life time.Duration) error {
	d.Lock()
	defer d.Unlock()
	t := new(Cache)
	t.createAt = time.Now()
	t.life = time.Duration(life) * time.Second
	t.funcs = funcs
	t.val = funcs()
	d.CachePools[key] = t

	return nil
}

//func test() interface{} {

//return rand.Intn(100)
//}

//func test1() interface{} {

//return rand.Float64()
//}
//func main() {
//cacheitem := NewCache()
//cacheitem.Put("test", test, 2)
//t := time.Now()
//for {
//m := cacheitem.Get("test")
//if m == nil {
//break
//}
//fmt.Println(m)
//}
//elapsed := time.Since(t)
//fmt.Println("app elapsed:", elapsed)
//}
