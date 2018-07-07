package cache

/**
 * DESCRIPTION: 实现lru cache功能
 *
 * @author rd
 * @create 2018-06-29 11:43
 **/
type lruCache struct {
    getValue func(key interface{}) (interface{}, error) //如果某个key寻找的对象不存在时，通过该方法获取需要的值
    compare  func(key, value interface{}) bool          //某个对象和所寻找的key是否相同
    size     int                                        //已有长度
    cap      int                                        //容量
    data     *link                                      //用户数据，带有头指针的链biao
}
type link struct {
    data interface{}
    next *link
}

func (l *lruCache) add(value interface{}, secondLast *link) {
    if l.size < l.cap {
        data := &link{
            data: value,
            next: l.data.next,
        }
        l.data.next = data
        l.size += 1
        return
    }
    if secondLast == nil {
        secondLast = l.data
        for i := 0; i < l.size; i++ {
            if secondLast.next.next == nil {
                break
            }
            secondLast = secondLast.next
        }
    }

    secondLast.next.data = value
    secondLast.next.next = l.data.next
    l.data.next = secondLast.next
    secondLast.next = nil
}

func (l *lruCache) Get(key interface{}) (interface{}, error) {
    if l.data.next == nil {
        res, err := l.getValue(key)
        if err != nil {
            return nil, err
        }
        l.add(res, nil)
        return res, nil
    }

    secondLast := l.data
    for {
        if l.compare(key, secondLast.next.data) {
            last := secondLast.next
            secondLast.next = last.next
            last.next = l.data.next
            l.data.next = last
            return l.data.next.data, nil
        }
        if secondLast.next.next == nil {
            break
        }
        secondLast = secondLast.next
    }

    res, err := l.getValue(key)
    if err != nil {
        return nil, err
    }
    l.add(res, secondLast)
    return res, nil
}

// 创建lru缓存
// num 缓存数量
// f 当获取数据从缓存获取不到时，f会从实际存储地获取数据
func NewLruCache(num int, getValue func(key interface{}) (interface{}, error), compare func(key, value interface{}) bool) Cache {
    return &lruCache{
        getValue: getValue,
        compare:  compare,
        cap:      num,
        data:     &link{},
    }
}
