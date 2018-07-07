package cache

import (
    "testing"
    "fmt"
    "strings"
    "time"
)

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-01 11:25
 **/

func TestNewLruCache(t *testing.T) {
    getValue := func(key interface{}) (interface{}, error) {
        time.Sleep(time.Microsecond * 100)
        switch k := key.(type) {
        case string:
            return k + "_value", nil
        case int:
            return k + 100, nil
        default:
            return nil, fmt.Errorf("unknown key")
        }
    }

    dis := func(l *link) []interface{} {
        items := []interface{}{}
        for l.next != nil {
            items = append(items, l.next.data)
            l = l.next
        }

        return items
    }

    compare := func(key, value interface{}) bool {
        switch k := key.(type) {
        case string:
            v, ok := value.(string)
            if !ok {
                return false
            }
            return strings.EqualFold(v, k+"_value")
        case int:
            v, ok := value.(int)
            if !ok {
                return false
            }
            return v == k+100
        }
        return false

    }

    cache := NewLruCache(5, getValue, compare)
    lruCache := cache.(*lruCache)

    params := []struct {
        Expected []interface{}
        In       interface{}
    }{
        {[]interface{}{"a_value"}, "a",},
        {[]interface{}{"b_value", "a_value"}, "b",},
        {[]interface{}{"c_value", "b_value", "a_value"}, "c"},
        {[]interface{}{103, "c_value", "b_value", "a_value"}, 3},
        {[]interface{}{"d_value", 103, "c_value", "b_value", "a_value"}, "d"},
        {[]interface{}{"a_value", "d_value", 103, "c_value", "b_value"}, "a"},
        {[]interface{}{"a_value", "d_value", 103, "c_value", "b_value"}, "a"},
        {[]interface{}{103, "a_value", "d_value", "c_value", "b_value"}, 3},
        {[]interface{}{"e_value", 103, "a_value", "d_value", "c_value"}, "e"},
    }
    for _, param := range params {
        res, err := lruCache.Get(param.In)
        if err != nil {
            t.Errorf("get errors:%s", err)
        }
        if res != param.Expected[0] {
            t.Errorf("expected:%+v, actual:%+v", param.Expected[0], res)
        }
        items := dis(lruCache.data)
        for i := range items {
            if items[i] != param.Expected[i] {
                t.Errorf("expected:%+v, actuval:%+v", param.Expected, items)
            }
        }
    }
    _, err := lruCache.Get(3.123)
    if !strings.EqualFold(err.Error(), "unknown key") {
        t.Errorf("expected:%+v, actuval:%+v", "unknow key", err)
    }
}
