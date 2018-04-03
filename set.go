package common

const (
    INITIAL_CAPACITY = 16
    LOADING_FACTOR   = 0.75
    BASE_NUMBER      = 31
)

type HashInt struct {
    l    int
    data [][]int
    cap  int
}

//是否包含某元素，包含返回true, 不包含返回false
func (h *HashInt) Contains(v int) bool {
    if h.data == nil || h.l == 0 {
        return false
    }
    p := (v * BASE_NUMBER) % h.cap
    items := h.data[p]
    for _, t := range items {
        if t == v {
            return true
        }
    }
    return false
}

//为hash新增某个元素
func (h *HashInt) Add(v int) {
    if len(h.data) == 0 {
        *h = HashInt{
            data: make([][]int, INITIAL_CAPACITY),
            cap:  INITIAL_CAPACITY,
        }
    }
    if h.Contains(v) {
        return
    }
    p := (v * BASE_NUMBER) % h.cap
    h.data[p] = append(h.data[p], v)
    h.l += 1
    if h.l >= int(float64(h.cap)*LOADING_FACTOR) {
        h.expansion()
    }
}

//进行扩容
func (h *HashInt) expansion() {
    tmp := &HashInt{
        data: make([][]int, h.cap*2),
        cap:  h.cap * 2,
    }
    for _, items := range h.data {
        if len(items) != 0 {
            for _, item := range items {
                tmp.Add(item)
            }
        }
    }
    *h = *tmp
}

//移除元素
func (h *HashInt) Remove(v int) {
    if len(h.data) == 0 {
        return
    }
    p := (v * BASE_NUMBER) % h.cap
    items := h.data[p]
    for i, t := range items {
        if t == v {
            items = append(items[:i], items[i+1:]...)
            break
        }
    }
    h.data[p] = items
}

//返回长度
func (h *HashInt) Len() int {
    if h == nil {
        return 0
    }
    return h.l
}

//返回容量
func (h *HashInt) Cap() int {
    if h == nil {
        return 0
    }
    return h.cap
}

func (h *HashInt) ForEach(f func(v int)) {
    if h.Len() == 0 {
        return
    }
    for _, items := range h.data {
        for _, item := range items {
            f(item)
        }
    }
}
