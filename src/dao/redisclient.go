/**导入redis包时，init初始化，生成redis池全局变量RedisPoolOne，
 *并在redis池中加入商品被秒件数“0”
 *封装了get\set\Lrang\Rpush等接口
 */
package dao

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	"vo"
)

var RedisPoolOne *RedisClient

func init() {
	RedisPoolOne = GetRedisInstance()
	RedisPoolOne.Set(vo.Product_Pre+vo.Product1_Query_Name, "0")
	RedisPoolOne.Set(vo.Product_Pre+vo.Product2_Query_Name, "0")
	RedisPoolOne.Set(vo.Product_Pre+vo.Product3_Query_Name, "0")
}

//定义redisPool对象连接ip、超时时长等属性
type RedisClient struct {
	pool         *redis.Pool
	redissvr     string
	conntimeout  int
	readtimeout  int
	writetimeout int
	maxidle      int
	maxactive    int
	expiresecond int
}

func GetRedisInstance() *RedisClient {
	redissvr := vo.Ip + ":" + vo.Port
	conntimeout := 100
	readtimeout := 50
	writetimeout := 50
	maxidle := 500
	maxactive := 1000
	expiresecond := 7000
	rc := new(RedisClient)
	if rc == nil {
		return nil
	}

	rc.pool = GetRedisPool(redissvr, conntimeout, readtimeout, writetimeout, maxidle, maxactive)
	if rc.pool == nil {
		return nil
	}

	rc.redissvr = redissvr

	rc.conntimeout = conntimeout
	rc.readtimeout = readtimeout
	rc.writetimeout = writetimeout
	rc.maxidle = maxidle
	rc.maxactive = maxactive
	rc.expiresecond = expiresecond

	return rc
}

func (rc *RedisClient) Exists(key string) (bool, error) {
	c := rc.pool.Get()
	defer c.Close()
	exists, err := redis.Bool((c.Do("EXISTS", key)))
	if err != nil {
		return false, err // handle error return from c.Do or type conversion error.
	}
	return exists, err
}
func (rc *RedisClient) Set(key, value string) error {
	c := rc.pool.Get()
	defer c.Close()

	reply, err := redis.String((c.Do("SET", key, value)))
	if err != nil {
		return err
	}

	// add redis key expire time.
	// ignore if error of expire command.
	rc.Expire(key, rc.expiresecond)

	if reply == "OK" {
		return nil
	} else {
		return errors.New("redisclient: unexpected reply of set")
	}
}

func (rc *RedisClient) Get(key string) (string, error) {
	c := rc.pool.Get()
	defer c.Close()

	reply, err := redis.String(c.Do("GET", key))
	if err != nil {
		return "", err
	}
	return reply, nil
}

func (rc *RedisClient) Expire(key string, expiresecond int) error {
	c := rc.pool.Get()
	defer c.Close()

	reply, err := redis.Int(c.Do("EXPIRE", key, expiresecond))
	if err != nil {
		return err
	}

	if reply == 1 {
		return nil
	} else {
		return errors.New("redisclient: unexpected reply of expire")
	}
}

func (rc *RedisClient) Del(key string) error {
	c := rc.pool.Get()
	defer c.Close()

	_, err := redis.Int(c.Do("DEL", key))
	if err != nil {
		return err
	}
	//	if reply == 1 {
	//		return nil
	//	} else {
	// reply为0时说明key不存在
	//		return errors.New("redisclient: unexpected reply of del")
	//	}
	return nil
}

func (rc *RedisClient) HGetall(key string) (map[string]string, error) {
	c := rc.pool.Get()
	defer c.Close()

	reply, err := redis.StringMap(c.Do("HGETALL", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (rc *RedisClient) HGet(key, subkey string) (string, error) {
	c := rc.pool.Get()
	defer c.Close()

	reply, err := redis.String(c.Do("HGET", key, subkey))
	if err != nil {
		return "", err
	}

	return reply, nil
}

func (rc *RedisClient) HSet(key, subkey, value string) error {
	c := rc.pool.Get()
	defer c.Close()

	_, err := redis.Int(c.Do("HSET", key, subkey, value))
	if err != nil {
		return err
	}

	// add redis key expire time.
	// ignore if error of expire command.
	rc.Expire(key, rc.expiresecond)

	// no need to check reply of HSET
	// reply == 1 means HSET key subkey value, subkey not exist
	// reply == 0 means HSET key subkey value, subkey exists, but the value is already modified.
	/*
		if reply == 1 {
			return nil
		} else {
			return errors.New("redisclient: unexpected reply of hset")
		}
	*/

	return nil
}

func (rc *RedisClient) RPush(key, value string) error {
	c := rc.pool.Get()
	defer c.Close()

	_, err := redis.Int(c.Do("rpush", key, value))
	if err != nil {
		return err
	}

	// add redis key expire time.
	// ignore if error of expire command.
	rc.Expire(key, rc.expiresecond)

	// no need to check reply of HSET
	// reply == 1 means HSET key subkey value, subkey not exist
	// reply == 0 means HSET key subkey value, subkey exists, but the value is already modified.
	/*
		if reply == 1 {
			return nil
		} else {
			return errors.New("redisclient: unexpected reply of hset")
		}
	*/

	return nil
}

func (rc *RedisClient) RPop(key string) (string, error) {
	c := rc.pool.Get()
	defer c.Close()

	reply, err := redis.String(c.Do("LPop", key))
	if err != nil {
		return "", err
	}
	return reply, nil
}

func (rc *RedisClient) LRange(key string) ([]string, error) {
	c := rc.pool.Get()
	defer c.Close()

	reply, err := redis.Strings(c.Do("LRANGE", key, 0, -1))
	if err != nil {
		return nil, err
	}
	return reply, nil
}
