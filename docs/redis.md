## redis
#### 1.简介
+ Redis支持数据的持久化，可以将内存中的数据保存在磁盘中，重启的时候可以再次加载使用。
+ Redis不仅支持简单的key-value类型的数据，同时还提供了list、set、zset、hash等数据结构的存储。
+ Redis支持数据的备份，即master-slave模式的数据备份。

#### 2. 打开客户端
+ 打开Redis客户端
```bash
redis-cli

redis-cli -h host -p port -a password
```

#### 3. Redis键（key）
相关命令：
```bash
SET runoobkey redis       # key=runoobkey，value=redis
GET runoobkey             # redis
DEL key                   # key存在的话删除key
EXISTS key                # 检查key是否存在
DUMP key                  # 序列化给定key，并返回序列化的值
EXPIRE key seconds        # 设置过期时间
EXPIREAT key timestamp
PEXPIRE key milliseconds  
PEXPIRE key milliseconds-timestamp
KEYS pattern
MOVE key db
PERSIST key
PTTL key
TTL key
RANDOMKEY
RENAME key newkey
RENAMENX key newkey
SCAN cursor [MATCH pattern] [COUNT count]
TYPE key
```

#### 4. Redis字符串（String）
value为string
```bash
SET key value
GET key
GETRANGE key start end
GETSET key value
GETBIT key offset
MGET key1 [key2...]
SETBIT key offset value
....
```

#### 5. Redis哈希（Hash）
value为哈希表
```bash
HMSET rbkey name "redis tutorial" description "redis basic commands" likes 20 visitors 23000

HDEL key field1 [field2]         # 删除哈希表中的一个字段
HGET key field                   # 获取哈希表中的一个字段
HGETALL key
HINCRBY key field increment
HKEYS key
HMGET key field1 [field2]
HMSET key field1 value1 [field2 value2]
HSET key field value
HSETNX key field value
HVALS key
HSCAN key cursor [MATCH pattern] [COUNT count]
```

#### 6. Redis列表
```bash
LPUSH runoobkey redis
LRANGE runoobkey 0 10
BLPOP key1 [key2] timeout
BRPOP key1 [key2] timeout
LINDEX key index
LINSERT key BEFORE|AFTER pivot value
LLEN key
LPOP key
LPUSH key value1 [value2]
LPUSHX key value
LRANGE key start stop
LREM key count value
LSET key index value
LTRIM key start stop
RPOP key
RPOPLPUSH source destination
RPUSH key value1 [value2]
RPUSHX key value
```

#### 7. Redis集合（SET）
无序集合，元素不能重复。
```bash
SADD key value
SMEMBERS key
SCARD key
SDIFF key1 [key2]
SDIFFSTORE destination key1 [key2]
SINTER key1 [key2]
SINTERSTORE destination key1 [key2]
SISMEMBER key member
SMEMBERS key
SMOVE source destination member
SPOP key
SRANDMEMBER key [count]
SREM key member1 [member2]
SUNIONSTORE destination key1 [key2]
SSCAN key cursor [MATCH pattern] [COUNT count]
```

#### 8. Redis有序集合（sorted set）
有序集合和集合一样也是string类型元素的集合，且不允许重复的成员。不同的是每个元素都会关联一个double类型的分数，redis这是通过分数来为集合中的成员进行从小到大的排序。
```bash
ZADD runoobkey 1 redis
ZRANGE runoobkey 0 10 WITHSCORES

ZCARD key              # 获取成员数
ZCOUNT key min max     # 分数在min、max之间的成员数
ZINCRBY key increment member
ZINTERSTORE destination numkeys key [key ...]
ZLECOUNT key min max
ZRANGE key start stop [WITHSCORES]
ZRANGEBYLEX key min max [LIMIT offset count]
ZRANGESCORE key min max [WITHSCORES] [LIMIT]
ZRANK key member
ZREM key member[member...]
....

```