package common

import (
    "testing"
)

func TestHashInt(t *testing.T) {
    hash := &HashInt{}
    hash.Add(1)
    tmp := map[int]bool{}
    for i := 0; i < 50; i++ {
        hash.Add(i)
        tmp[i] = true
    }
    for i := 0; i < hash.Cap(); i++ {
        if tmp[i] != hash.Contains(i) {
            t.Errorf("invalid ")
        }
    }
    hash.Remove(3)
    if hash.Contains(3) {
        t.Errorf("not contains 3")
    }

}
