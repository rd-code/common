package filter

import (
    "testing"
    "fmt"
    "strings"
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
    params := []Param{
        {"a", true},
        {"b", false},
        {"c", true},
    }
    items := Filter(params, func(index int) bool {
        return params[index].value
    })
    res := []Param{
        {"a", true},
        {"c", true},
    }
    if len(res) != len(items) {
        t.Errorf("expected size:%d, actual size:%s", len(res), len(items))
    }
    for i, v := range res {
        item, ok := items[i].(Param)
        if !ok {
            t.Fatalf("element cannot convert to Param")
        }
        if !strings.EqualFold(item.data, v.data) || item.value != v.value {
            t.Fatalf("exptected element:%+v, actuval element:%+v", v, item)
        }
    }
    fmt.Println(items)
}
