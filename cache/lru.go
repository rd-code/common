package cache

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
    keys *lruCache //第一个为头指针
}

func (l *lruCache) add(key string, value interface{}) {

}

func (l *lruCache) Get(key string) interface{} {
    return nil
}

type link struct {
    data string
    next *link
}

func NewLruCache(num int, f func(key string) interface{}) Cache {
    return &lruCache{
        f:    f,
        cap:  num,
        data: make(map[string]interface{}, num),
        keys: &lruCache{},
    }
}
