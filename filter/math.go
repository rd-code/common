package filter

import "reflect"

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-04 17:32
 **/

//从指定slice获取最大整数
//data为一个slice
//f 通过下表获取对应的整数
func MaxInt(v interface{}, f func(index int) int) int {
    rv := reflect.ValueOf(v)
    if rv.Len() == 0 {
        return 0
    }

    res := f(0)
    for i := 0; i < rv.Len(); i++ {
        value := f(i)
        if value > res {
            res = value
        }
    }
    return res
}

//同MaxInt
func MaxInt64(v interface{}, f func(index int) int64) int64 {
    rv := reflect.ValueOf(v)
    if rv.Len() == 0 {
        return 0
    }

    res := f(0)
    for i := 0; i < rv.Len(); i++ {
        value := f(i)
        if value > res {
            res = value
        }
    }
    return res
}

//同maxInt
func MaxInt32(v interface{}, f func(index int) int32) int32 {
    rv := reflect.ValueOf(v)
    if rv.Len() == 0 {
        return 0
    }

    res := f(0)
    for i := 0; i < rv.Len(); i++ {
        value := f(i)
        if value > res {
            res = value
        }
    }
    return res
}

//从指定slice获取最小值
//data为一个slice
//f 通过下表获取对应的整数
func MinInt(v interface{}, f func(index int) int) int {
    rv := reflect.ValueOf(v)
    if rv.Len() == 0 {
        return 0
    }

    res := f(0)
    for i := 0; i < rv.Len(); i++ {
        value := f(i)
        if value < res {
            res = value
        }
    }
    return res
}

//同minInt
func MinInt64(v interface{}, f func(index int) int64) int64 {
    rv := reflect.ValueOf(v)
    if rv.Len() == 0 {
        return 0
    }

    res := f(0)
    for i := 0; i < rv.Len(); i++ {
        value := f(i)
        if value < res {
            res = value
        }
    }
    return res
}

//同minInt
func MinInt32(v interface{}, f func(index int) int32) int32 {
    rv := reflect.ValueOf(v)
    if rv.Len() == 0 {
        return 0
    }

    res := f(0)
    for i := 0; i < rv.Len(); i++ {
        value := f(i)
        if value < res {
            res = value
        }
    }
    return res
}
