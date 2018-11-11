package filter

import (
    "fmt"
    "reflect"
)

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-04 17:42
 **/

//Filter 从slice中过滤信息到slice
//src 源slice
//dst 目的slice,需要为slice指针
//f 过滤函数返回value, ok，当ok为true时，需要将value加到dst上
//type Param struct {
//        data  string
//        value bool
//    }
//    src := []Param{
//        {"a", true},
//        {"b", false},
//        {"c", true},
//    }
//    var dst []string
//    Filter(src, &dst, func(index int) (interface{}, bool) {
//        return src[index].data, src[index].value
//    })
func Filter(src, dst interface{}, f func(index int) (interface{}, bool)) {
    rv := reflect.ValueOf(src)
    items := make([]interface{}, 0, rv.Len())
    for i := 0; i < rv.Len(); i++ {
        value, ok := f(i)
        if ok {
            items = append(items, value)
        }
    }
    slice := reflect.MakeSlice(reflect.TypeOf(dst).Elem(), len(items), len(items))
    for i, item := range items {
        fmt.Println(slice.Index(i).Kind())
        slice.Index(i).Set(reflect.ValueOf(item))
    }
    reflect.ValueOf(dst).Elem().Set(slice)
}
