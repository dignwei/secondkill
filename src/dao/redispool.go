/**提供生成redisPool的接口
 *进行连接错误处理
 */
package dao

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	"sync"
	"time"
)

var redispool_init_ctx sync.Once
var redispool_instance *redis.Pool

func GetRedisPool(redissvr string, conntimeout, readtimeout, writetimeout, maxidle, maxactive int) *redis.Pool {

	redispool_init_ctx.Do(func() {

		redispool_instance = &redis.Pool{
			MaxIdle:   maxidle,
			MaxActive: maxactive,
			Dial: func() (redis.Conn, error) {

				c, err := redis.DialTimeout("tcp", redissvr, time.Duration(conntimeout)*time.Millisecond, time.Duration(readtimeout)*time.Millisecond, time.Duration(writetimeout)*time.Millisecond)
				if err == nil && c != nil {
					return c, nil
				}

				return nil, errors.New("redispool: cannot connect to any redis server")
			},
		}
	})

	return redispool_instance
}
