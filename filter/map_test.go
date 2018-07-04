package filter

import (
    "testing"
    "strconv"
    "strings"
)

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-01 14:50
 **/

type data struct {
    name string
    age  int
    sex  bool
}

func TestMap(t *testing.T) {
    items := make([]data, 0)
    for i := 0; i < 10; i++ {
        items = append(items, data{
            name: "name" + strconv.Itoa(i),
            age:  20 + i,
            sex:  i%2 == 0,
        })
    }

    names := MapString(items, func(index int) string {
        return items[index].name
    })

    ages := MapInt(items, func(index int) int {
        return items[index].age
    })

    for i, name := range names {
        if !strings.EqualFold(name, "name"+strconv.Itoa(i)) {
            t.Error("test failed")
        }
    }

    for i, age := range ages {
        if i != (age - 20) {
            t.Error("test failed")
        }
    }
}
