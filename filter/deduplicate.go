package filter

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-12 13:59
 **/

//请使用DeduplicateString
func RemoveString(items []string) []string {
    return DeduplicateString(items)
}

//删除重复数据
//返回一个去重之后的string slice
func DeduplicateString(items []string) []string {
    data := make(map[string]struct{}, len(items))
    for _, item := range items {
        data[item] = struct{}{}
    }
    res := make([]string, 0, len(data))
    for k := range data {
        res = append(res, k)
    }
    return res
}

//同DeduplicateString
func DeduplicateInt(items []int) []int {
    data := make(map[int]struct{}, len(items))
    for _, item := range items {
        data[item] = struct{}{}
    }
    res := make([]int, 0, len(data))
    for k := range data {
        res = append(res, k)
    }
    return res
}

//同DeduplicateString
func DeduplicateInt64(items []int64) []int64 {
    data := make(map[int64]struct{}, len(items))
    for _, item := range items {
        data[item] = struct{}{}
    }
    res := make([]int64, 0, len(data))
    for k := range data {
        res = append(res, k)
    }
    return res
}

//同DeduplicateString
func DeduplicateInt32(items []int32) []int32 {
    data := make(map[int32]struct{}, len(items))
    for _, item := range items {
        data[item] = struct{}{}
    }
    res := make([]int32, 0, len(data))
    for k := range data {
        res = append(res, k)
    }
    return res
}

//同DeduplicateString
func DeduplicateInt16(items []int16) []int16 {
    data := make(map[int16]struct{}, len(items))
    for _, item := range items {
        data[item] = struct{}{}
    }
    res := make([]int16, 0, len(data))
    for k := range data {
        res = append(res, k)
    }
    return res
}

//同DeduplicateString
func DeduplicateInt8(items []int8) []int8 {
    data := make(map[int8]struct{}, len(items))
    for _, item := range items {
        data[item] = struct{}{}
    }
    res := make([]int8, 0, len(data))
    for k := range data {
        res = append(res, k)
    }
    return res
}

//同DeduplicateString
func DeduplicateUint(items []uint) []uint {
    data := make(map[uint]struct{}, len(items))
    for _, item := range items {
        data[item] = struct{}{}
    }
    res := make([]uint, 0, len(data))
    for k := range data {
        res = append(res, k)
    }
    return res
}

//同DeduplicateString
func DeduplicateUint64(items []uint64) []uint64 {
    data := make(map[uint64]struct{}, len(items))
    for _, item := range items {
        data[item] = struct{}{}
    }
    res := make([]uint64, 0, len(data))
    for k := range data {
        res = append(res, k)
    }
    return res
}

//同DeduplicateString
func DeduplicateUint32(items []uint32) []uint32 {
    data := make(map[uint32]struct{}, len(items))
    for _, item := range items {
        data[item] = struct{}{}
    }
    res := make([]uint32, 0, len(data))
    for k := range data {
        res = append(res, k)
    }
    return res
}

//同DeduplicateString
func DeduplicateUint16(items []uint16) []uint16 {
    data := make(map[uint16]struct{}, len(items))
    for _, item := range items {
        data[item] = struct{}{}
    }
    res := make([]uint16, 0, len(data))
    for k := range data {
        res = append(res, k)
    }
    return res
}

//同DeduplicateString
func DeduplicateUint8(items []uint8) []uint8 {
    data := make(map[uint8]struct{}, len(items))
    for _, item := range items {
        data[item] = struct{}{}
    }
    res := make([]uint8, 0, len(data))
    for k := range data {
        res = append(res, k)
    }
    return res
}

//同DeduplicateString
func DeduplicateBool(items []bool) []bool {
    data := make(map[bool]struct{}, len(items))
    for _, item := range items {
        data[item] = struct{}{}
    }
    res := make([]bool, 0, len(data))
    for k := range data {
        res = append(res, k)
    }
    return res
}

//同DeduplicateString
func DeduplicateFloat64(items []float64) []float64 {
    data := make(map[float64]struct{}, len(items))
    for _, item := range items {
        data[item] = struct{}{}
    }
    res := make([]float64, 0, len(data))
    for k := range data {
        res = append(res, k)
    }
    return res
}

//同DeduplicateString
func DeduplicateFloat32(items []float32) []float32 {
    data := make(map[float32]struct{}, len(items))
    for _, item := range items {
        data[item] = struct{}{}
    }
    res := make([]float32, 0, len(data))
    for k := range data {
        res = append(res, k)
    }
    return res
}

//同DeduplicateString
func DeduplicateComplex128(items []complex128) []complex128 {
    data := make(map[complex128]struct{}, len(items))
    for _, item := range items {
        data[item] = struct{}{}
    }
    res := make([]complex128, 0, len(data))
    for k := range data {
        res = append(res, k)
    }
    return res
}

//同DeduplicateString
func DeduplicateComplex64(items []complex64) []complex64 {
    data := make(map[complex64]struct{}, len(items))
    for _, item := range items {
        data[item] = struct{}{}
    }
    res := make([]complex64, 0, len(data))
    for k := range data {
        res = append(res, k)
    }
    return res
}
