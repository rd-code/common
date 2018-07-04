package math

import (
    "testing"
    "strings"
)

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-04 17:02
 **/

func TestFormatInt(t *testing.T) {
    type Param struct {
        expected string
        data     int64
        base     int
    }

    params := []Param{
        {"0", 0, 3},
        {"0", 0, 27},
        {"1", 1, 9},
        {"1", 1, 15},
        {"10", 4, 4},
        {"10", 62, 62},
        {"10", 27, 27},
        {"11", 63, 62},
        {"11", 6, 5},
        {"11", 17, 16},
        {"20", 124, 62},
    }

    for _, param := range params {
        if !strings.EqualFold(param.expected, FormatInt(param.data, param.base)) {
            t.Errorf("exptected value:%s, actual value:%s, data:%d, base:%d",
                param.expected, FormatInt(param.data, param.base), param.data, param.base)
        }
    }
}
