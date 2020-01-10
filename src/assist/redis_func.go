package assist

/*
 * redis proxy for logic modules.
 * author: CC
 * email : 151503324@qq.com
 * date  : 2017.06.17
 */

import (
	"cherry/base"
	"cherry/dbproxy"
	"errors"

	"github.com/garyburd/redigo/redis"
)

// RedisExists 返回值1: -1表示错误, 0表示key不存在, 1表示key存在.
func RedisExists(insRedis *dbproxy.Redis, key interface{}) (int32, error) {
	reply, err := redis.Int(insRedis.ReadCommand("EXISTS", key))
	if err != nil {
		base.LogError("RedisExists:", err)
		return -1, errors.New(err.Error())
	}
	if reply == 0 {
		return 0, nil
	}
	return 1, errors.New("RedisExists the key is exists")
}

// RedisINCR TODO
func RedisINCR(insRedis *dbproxy.Redis, key interface{}) (uint64, error) {
	reply, err := redis.Uint64(insRedis.WriteCommand("INCR", key))
	if err != nil {
		base.LogError("RedisINCR:", err)
		return 0, err
	}
	if reply == 0 {
		return 0, errors.New("RedisINCR data error")
	}
	return reply, nil
}

// RedisGET TODO
func RedisGET(insRedis *dbproxy.Redis, key interface{}) (interface{}, error) {
	reply, err := insRedis.ReadCommand("GET", key)
	if reply == nil || err != nil {
		base.LogError("RedisGET:", err)
		return nil, err
	}
	return reply, nil
}

// RedisSET TODO
func RedisSET(insRedis *dbproxy.Redis, key interface{}, value interface{}) error {
	_, err := insRedis.WriteCommand("SET", key, dbproxy.RedisMarshal(value))
	if err != nil {
		base.LogError("RedisSET:", err)
		return errors.New(err.Error())
	}
	return nil
}

// RedisDEL TODO
func RedisDEL(insRedis *dbproxy.Redis, key interface{}) error {
	_, err := insRedis.WriteCommand("DEL", key)
	if err != nil {
		base.LogError("RedisDEL:", err)
		return errors.New(err.Error())
	}
	return nil
}

// // MGET.
// func RedisMGET(insRedis *dbproxy.Redis, args ...interface{}) {
// }

// // MSET.
// func RedisMSET(insRedis *dbproxy.Redis, args ...interface{}) {
// }

///////////////////////////////////////////////////////////////
// Redis Hash.

// RedisHEXISTS 返回值1: -1表示错误, 0表示key不存在, 1表示key存在.
func RedisHEXISTS(insRedis *dbproxy.Redis, key, field interface{}) (int32, error) {
	reply, err := redis.Int(insRedis.ReadCommand("HEXISTS", key, field))
	if err != nil {
		base.LogError("RedisHEXISTS:", err)
		return -1, errors.New(err.Error())
	}
	if reply == 0 {
		return 0, nil
	}
	return 1, errors.New("RedisExists the key is exists")
}

// // HLEN.
// func RedisHLEN(insRedis *dbproxy.Redis, key interface{}) uint64 {
// 	return 0
// }

// RedisHVALS 需要对数组遍历使用时通过 dbproxy.RedisUnmarshal进行类型转换.
func RedisHVALS(insRedis *dbproxy.Redis, key interface{}) ([]interface{}, error) {
	reply, err := redis.Values(insRedis.ReadCommand("HVALS", key))
	if reply == nil || err != nil {
		base.LogError("RedisHVALS:", err)
		return nil, err
	}
	return reply, nil
}

// HGETALL.

// RedisHGET TODO
func RedisHGET(insRedis *dbproxy.Redis, key, field interface{}) (interface{}, error) {
	reply, err := insRedis.ReadCommand("HGET", key, field)
	if reply == nil || err != nil {
		base.LogError("RedisHGET:", err)
		return nil, err
	}
	return reply, nil
}

// RedisHSET TODO
func RedisHSET(insRedis *dbproxy.Redis, key, field, value interface{}) error {
	_, err := insRedis.WriteCommand("HSET", key, field, dbproxy.RedisMarshal(value))
	if err != nil {
		base.LogError("RedisHSET:", err)
		return errors.New(err.Error())
	}
	return nil
}

// RedisHDEL TODO
func RedisHDEL(insRedis *dbproxy.Redis, key, field interface{}) error {
	_, err := insRedis.WriteCommand("HDEL", key, field)
	if err != nil {
		base.LogError("RedisHDEL:", err)
		return errors.New(err.Error())
	}
	return nil
}

// // HMGET.
// func RedisHMGET(insRedis *dbproxy.Redis, key interface{}, fields ...interface{}) interface{} {
// 	return nil
// }

// // HMSET.
// func RedisHMSET(insRedis *dbproxy.Redis, key interface{}, args ...interface{}) bool {
// 	return false
// }
