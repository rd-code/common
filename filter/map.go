package filter

import (
    "reflect"
)

/**
 * DESCRIPTION: 完成某些过滤功能
 *
 * @author rd
 * @create 2018-07-01 14:39
 **/
//Map 从源slice的元素里面挑选信息到目的slice里面
//src 源slice
//dst 目的slice
//f 挑选函数
func Map(src interface{}, dst interface{}, f func(index int) interface{}) {
    srcRV := reflect.ValueOf(src)
    slice := reflect.MakeSlice(reflect.TypeOf(dst).Elem(), srcRV.Len(), srcRV.Len())

    for i := 0; i < srcRV.Len(); i++ {
        value := f(i)
        slice.Index(i).Set(reflect.ValueOf(value))
    }
    reflect.ValueOf(dst).Elem().Set(slice)

}

//同Map 返回为[]string
func MapString(v interface{}, f func(index int) string) ([]string) {
    rv := reflect.ValueOf(v)
    res := make([]string, rv.Len())
    for i := 0; i < rv.Len(); i++ {
        res[i] = f(i)
    }
    return res
}

//同Map 返回为[]int
func MapInt(v interface{}, f func(index int) int) ([]int) {
    rv := reflect.ValueOf(v)
    res := make([]int, rv.Len())
    for i := 0; i < rv.Len(); i++ {
        res[i] = f(i)
    }
    return res
}

//同Map 返回为[]int32
func MapInt32(v interface{}, f func(index int) int32) []int32 {
    rv := reflect.ValueOf(v)
    res := make([]int32, rv.Len())
    for i := 0; i < rv.Len(); i++ {
        res[i] = f(i)
    }
    return res
}

//同Map 返回为[]int64
func MapInt64(v interface{}, f func(index int) int64) []int64 {
    rv := reflect.ValueOf(v)
    res := make([]int64, rv.Len())
    for i := 0; i < rv.Len(); i++ {
        res[i] = f(i)
    }
    return res
}

//同Map 返回为[]bool
func MapBool(v interface{}, f func(index int) bool) []bool {
    rv := reflect.ValueOf(v)
    res := make([]bool, rv.Len())
    for i := 0; i < rv.Len(); i++ {
        res[i] = f(i)
    }
    return res
}

//同Map 返回为[]float32
func MapFloat32(v interface{}, f func(index int) float32) []float32 {
    rv := reflect.ValueOf(v)
    res := make([]float32, rv.Len())
    for i := 0; i < rv.Len(); i++ {
        res[i] = f(i)
    }
    return res
}

//同Map 返回为[]float64
func MapFloat64(v interface{}, f func(index int) float64) []float64 {
    rv := reflect.ValueOf(v)
    res := make([]float64, rv.Len())
    for i := 0; i < rv.Len(); i++ {
        res[i] = f(i)
    }
    return res
}
