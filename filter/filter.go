package filter

import "reflect"

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-04 17:42
 **/

//从指定slice过滤不满足需求slice
//data 源slice
//f 判断是否满足过滤需求
func Filter(data interface{}, f func(index int) bool) []interface{} {
    rv := reflect.ValueOf(data)
    res := make([]interface{}, 0, rv.Len())
    for i := 0; i < rv.Len(); i++ {
        if f(i) {
            res = append(res, rv.Index(i).Interface())
        }
    }
    return res
}
