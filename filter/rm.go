package filter

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-07-12 13:59
 **/

func RemoveString(items []string) []string {
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
