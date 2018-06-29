package cache

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-06-29 11:14
 **/

type Cache interface {
    Get(key string) interface{}
}
