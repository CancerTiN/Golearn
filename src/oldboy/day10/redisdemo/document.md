# Redis应用场景
- 缓存系统，减轻主数据库（MySQL）的压力
- 计数场景，比如微博、抖音中的关注数和粉丝数
- 热门排行榜，需要排序的场景特别适合使用ZSET
- 利用LIST可以实现队列的功能

# Redis与Memcached比较
- Redis支持5种数据类型，Memcached只支持字符串
- Redis的性能比Memcached好很多
- Redis支持RDB持久化和AOF持久化
- Redis支持master/slave模式
