package common

import (
    "testing"
    "fmt"
)

func TestStack(t *testing.T) {
    ss := NewStack()
    ss.Push("a")
    ss.Push("b")
    ss.Push("c")
    ss.Push("d")
    ss.Push("e")
    if ss.Len() != 5 {
        t.Errorf("invalid size ")
    }

    items := [][]interface{}{{"e", 4}, {"d", 3}, {"c", 2}, {"b", 1}, {"a", 0}}

    for !ss.IsEmpty() {
        tmp, err := ss.Pop()
        if err != nil {
            t.Error("pop failed \n", err)
        }
        if tmp != items[0][0] {
            t.Errorf("expected value:%v, actual value:%v", items[0][0], tmp)
        }
        if items[0][1] != ss.Len() {
            t.Errorf("exptected size:%v, actual size:%v", items[0][1], ss.Len())
        }
        items = items[1:]
    }
}
