package cache

import (
	"testing"
	//"time"
)

func TestCache(t *testing.T) {

	cacheitem := NewCache()

	cacheitem.Put("test", 456, 10)
	//m := time.NewTicker(time.Duration(time.Second * 1))
	//d := time.Duration(time.Second * 2)
	//m := time.NewTicker(d)
	t.Error(cacheitem.Get("test"))
	//defer m.Stop()
	for {

		t.Error(cacheitem.Get("test"))

	}
}
