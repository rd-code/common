package cache

import (
    "testing"
    "fmt"
    "strings"
)

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-01 11:25
 **/

func TestNewLruCache(t *testing.T) {
    f := func(key string) interface{} {
        return "value" + key
    }

    dis := func(l *link) {
        keys := []string{}
        for l.next != nil {
            keys = append(keys, l.next.data)
            l = l.next
        }
        fmt.Println(strings.Join(keys, ", "))

    }
    cache := NewLruCache(5, f)
    lruCache := cache.(*lruCache)

    cache.Get("a")
    dis(lruCache.keys)
    cache.Get("b")
    cache.Get("c")
    cache.Get("d")
    cache.Get("e")
    dis(lruCache.keys)
    cache.Get("a")
    dis(lruCache.keys)
    cache.Get("d")
    dis(lruCache.keys)
    cache.Get("f")
    dis(lruCache.keys)
}
