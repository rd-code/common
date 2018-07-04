package math

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-04 16:32
 **/

//将数字格式化成字符串
//数之间的大小关系为 0 < 9 < A < Z < a < z
//data 需要格式化的值，为一个数字
//base 进制 2到62
func FormatInt(data int64, base int) (string) {
    if base < 2 || base > 62 {
        panic("math: illegal FormatInt base")
    }
    if data < 0 {
        return "-" + formatInt(0-data, int64(base))
    }
    return formatInt(data, int64(base))
}

//每个字母的下表代表该字母的值
var strSlice = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
    "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N",
    "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
    "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n",
    "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
//将数字格式化成字符串
//不再进行参数校验以及特殊处理
func formatInt(data int64, base int64) (string) {
    res := ""
    for {
        if data < base {
            res = strSlice[data] + res
            break
        }
        mod := data % base
        res = strSlice[mod] + res
        data = data / base
        if data == 0 {
            break
        }
    }
    return res
}
