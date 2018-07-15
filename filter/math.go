package filter

import (
    "reflect"
)

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-04 17:32
 **/

//从一个整型slice里面获取最大的值
func MaxInt(items []int) int {
    var max int
    if len(items) > 0 {
        max = items[0]
    }
    for _, v := range items {
        if v > max {
            max = v
        }
    }
    return max
}

//参考maxInt
func MaxInt64(items []int64) int64 {
    var max int64
    if len(items) > 0 {
        max = items[0]
    }
    for _, v := range items {
        if v > max {
            max = v
        }
    }
    return max
}

//参考maxInt
func MaxInt32(items []int32) int32 {
    var max int32
    if len(items) > 0 {
        max = items[0]
    }
    for _, v := range items {
        if v > max {
            max = v
        }
    }
    return max
}

//参考maxInt
func MaxInt16(items []int16) int16 {
    var max int16
    if len(items) > 0 {
        max = items[0]
    }
    for _, v := range items {
        if v > max {
            max = v
        }
    }
    return max
}

//参考maxInt
func MaxInt8(items []int8) int8 {
    var max int8
    if len(items) > 0 {
        max = items[0]
    }
    for _, v := range items {
        if v > max {
            max = v
        }
    }
    return max
}

//参考maxInt
func MaxUint(items []uint) uint {
    var max uint
    if len(items) > 0 {
        max = items[0]
    }
    for _, v := range items {
        if v > max {
            max = v
        }
    }
    return max
}

//参考maxInt
func MaxUint64(items []uint64) uint64 {
    var max uint64
    if len(items) > 0 {
        max = items[0]
    }
    for _, v := range items {
        if v > max {
            max = v
        }
    }
    return max
}

//参考maxInt
func MaxUint32(items []uint32) uint32 {
    var max uint32
    if len(items) > 0 {
        max = items[0]
    }
    for _, v := range items {
        if v > max {
            max = v
        }
    }
    return max
}

//参考maxInt
func MaxUint16(items []uint16) uint16 {
    var max uint16
    if len(items) > 0 {
        max = items[0]
    }
    for _, v := range items {
        if v > max {
            max = v
        }
    }
    return max
}

//参考maxInt
func MaxUint8(items []uint8) uint8 {
    var max uint8
    if len(items) > 0 {
        max = items[0]
    }
    for _, v := range items {
        if v > max {
            max = v
        }
    }
    return max
}

//参考maxInt
func MaxFloat64(items []float64) float64 {
    var max float64
    if len(items) > 0 {
        max = items[0]
    }
    for _, v := range items {
        if v > max {
            max = v
        }
    }
    return max
}

//参考maxInt
func MaxFloat32(items []float32) float32 {
    var max float32
    if len(items) > 0 {
        max = items[0]
    }
    for _, v := range items {
        if v > max {
            max = v
        }
    }
    return max
}

//从一个整型slice里面获取最大的值
func MinInt(items []int) int {
    var min int
    if len(items) > 0 {
        min = items[0]
    }
    for _, v := range items {
        if v < min {
            min = v
        }
    }
    return min
}

//参考MinInt
func MinInt64(items []int64) int64 {
    var min int64
    if len(items) > 0 {
        min = items[0]
    }
    for _, v := range items {
        if v < min {
            min = v
        }
    }
    return min
}

//参考MinInt
func MinInt32(items []int32) int32 {
    var min int32
    if len(items) > 0 {
        min = items[0]
    }
    for _, v := range items {
        if v < min {
            min = v
        }
    }
    return min
}

//参考MinInt
func MinInt16(items []int16) int16 {
    var min int16
    if len(items) > 0 {
        min = items[0]
    }
    for _, v := range items {
        if v < min {
            min = v
        }
    }
    return min
}

//参考MinInt
func MinInt8(items []int8) int8 {
    var min int8
    if len(items) > 0 {
        min = items[0]
    }
    for _, v := range items {
        if v < min {
            min = v
        }
    }
    return min
}

//参考MinInt
func MinUint(items []uint) uint {
    var min uint
    if len(items) > 0 {
        min = items[0]
    }
    for _, v := range items {
        if v < min {
            min = v
        }
    }
    return min
}

//参考MinInt
func MinUint64(items []uint64) uint64 {
    var min uint64
    if len(items) > 0 {
        min = items[0]
    }
    for _, v := range items {
        if v < min {
            min = v
        }
    }
    return min
}

//参考MinInt
func MinUint32(items []uint32) uint32 {
    var min uint32
    if len(items) > 0 {
        min = items[0]
    }
    for _, v := range items {
        if v < min {
            min = v
        }
    }
    return min
}

//参考MinInt
func MinUint16(items []uint16) uint16 {
    var min uint16
    if len(items) > 0 {
        min = items[0]
    }
    for _, v := range items {
        if v < min {
            min = v
        }
    }
    return min
}

//参考MinInt
func MinUint8(items []uint8) uint8 {
    var min uint8
    if len(items) > 0 {
        min = items[0]
    }
    for _, v := range items {
        if v < min {
            min = v
        }
    }
    return min
}

//参考MinInt
func MinFloat64(items []float64) float64 {
    var min float64
    if len(items) > 0 {
        min = items[0]
    }
    for _, v := range items {
        if v < min {
            min = v
        }
    }
    return min
}

//参考MinInt
func MinFloat32(items []float32) float32 {
    var min float32
    if len(items) > 0 {
        min = items[0]
    }
    for _, v := range items {
        if v < min {
            min = v
        }
    }
    return min
}

//从指定slice查询最大的元素
//data 源slice
//greater 判断第一个参数是否大于第二个参数
func Max(data interface{}, greater func(v1, v2 interface{}) bool) interface{} {
    rv := reflect.ValueOf(data)
    var res, tmp interface{}
    if rv.Len() > 0 {
        res = rv.Index(0).Interface()
    }
    for i := 0; i < rv.Len(); i++ {
        tmp = rv.Index(i).Interface()
        if greater(tmp, res) {
            res = tmp
        }
    }
    return res
}

//从指定slice查询最大的元素
//data 源slice
//less 判断第一个参数是否小于第二个参数
func Min(data interface{}, less func(v1, v2 interface{}) bool) interface{} {
    rv := reflect.ValueOf(data)
    var res, tmp interface{}
    if rv.Len() > 0 {
        res = rv.Index(0).Interface()
    }
    for i := 0; i < rv.Len(); i++ {
        tmp = rv.Index(i).Interface()
        if less(tmp, res) {
            res = tmp
        }
    }
    return res
}
