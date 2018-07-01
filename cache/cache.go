package cache

/**
 * DESCRIPTION:
 *
 * @author rd
 * @create 2018-06-29 11:14
 **/
/**
实现缓存功能
 */
type Cache interface {
    Get(key string) interface{} //根据key值来获取缓存的值
}
