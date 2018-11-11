package filter

import (
    "strings"
    "testing"
)

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-04 17:46
 **/

func TestFilter(t *testing.T) {
    type Param struct {
        data  string
        value bool
    }
    src := []Param{
        {"a", true},
        {"b", false},
        {"c", true},
    }
    var dst []string
    Filter(src, &dst, func(index int) (interface{}, bool) {
        return src[index].data, src[index].value
    })

    res := []string{"a", "c",}
    if len(res) != len(dst) {
        t.Errorf("expected size:%d, actual size:%d", len(res), len(dst))
    }
    for i, v := range res {
        item := dst[i]
        if !strings.EqualFold(item, v) {
            t.Fatalf("exptected element:%+v, actuval element:%+v", v, item)
        }
    }
}
