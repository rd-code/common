package cache

import (
    "strings"
)

/**
 * DESCRIPTION: 实现lru cache功能
 *
 * @author rd
 * @create 2018-06-29 11:43
 **/
type lruCache struct {
    f    func(key string) interface{}
    size int
    cap  int
    data map[string]interface{}
    keys *link //第一个为头指针
}

type link struct {
    data string
    next *link
}

func (l *lruCache) add(key string, value interface{}) {
    l.data[key] = value
    if l.size < l.cap {
        data := &link{
            data: key,
            next: l.keys.next,
        }
        l.keys.next = data
        l.size += 1
        return
    }
    secondLast := l.keys
    for i := 0; i < l.size; i++ {
        if secondLast.next.next == nil {
            break
        }
        secondLast = secondLast.next
    }

    secondLast.next.data = key
    secondLast.next.next = l.keys.next
    l.keys.next = secondLast.next
    secondLast.next = nil
}

func (l *lruCache) Get(key string) interface{} {
    v, ok := l.data[key]
    if !ok {
        res := l.f(key)
        l.add(key, res)
        return res
    }

    prefix := l.keys
    for i := 0; i < l.size; i++ {
        if strings.EqualFold(prefix.next.data, key) {
            break
        }
        prefix = prefix.next
    }

    target := prefix.next

    prefix.next = target.next
    target.next = l.keys.next
    l.keys.next = target

    return v
}

// 创建lru缓存
// num 缓存数量
// f 当获取数据从缓存获取不到时，f会从实际存储地获取数据
func NewLruCache(num int, f func(key string) interface{}) Cache {
    return &lruCache{
        f:    f,
        cap:  num,
        data: make(map[string]interface{}, num),
        keys: &link{},
    }
}
