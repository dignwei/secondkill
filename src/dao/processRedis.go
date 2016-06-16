/**提供redis的直接接口
*用于循环携程listen——redis调用
 */
package dao

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strings"
)

var conn redis.Conn

//输入redis服务器ip与port,返回redis连接对象
func OpenRedis(IP string, Port string) {
	IPAndPort := IP + ":" + Port
	conn, _ = redis.Dial("tcp", IPAndPort)
	if conn == nil {
		fmt.Printf("redis连接失败\n")
	}
}

func CloseRedis() {
	conn.Close()
}

func Sadd(key string, value string) {
	conn.Do("SADD", key, value)
}

func SgetAll(key string) string {
	setMembers, _ := redis.Values(conn.Do("SMEMBERS", key))
	res := ""
	isFirst := true
	for _, itemV := range setMembers {
		tmpStr := fmt.Sprintf("%s", itemV)
		if isFirst {
			res += tmpStr
			isFirst = false
		} else {
			res += "|" + tmpStr
		}

	}
	return res
}

func Srem(key string, value string) {
	conn.Do("SREM", key, value)
}

func SisExist(key string, value string) bool {
	isEx, _ := redis.Bool(conn.Do("SISMEMBER", key, value))
	return isEx
}

func HSetValue(key string, value string) {
	conn.Do("SET", key, value)
}

//func INC(key string) {
//	conn.Do("INCR",key)
//}

func LPushValue(key string, value string) {
	_, err := conn.Do("lpush", key, value)
	if err != nil {
		fmt.Println("errMsg:", err)
	}
}

func LPopValue(key string) string {
	value, _ := redis.String(conn.Do("lpop", key))
	//if err != nil {
	//	fmt.Println("RPopValue:", err)
	//}
	return string(value)
}

func HGetValue(key string) string {
	value, _ := redis.String(conn.Do("GET", key))
	return string(value)
}
func HisExist(key string) bool {
	exists, _ := redis.Bool(conn.Do("EXISTS", key))
	return exists
}

func getTwoSetsUnion(set1 string, set2 string) string {
	setMembers, _ := redis.Values(conn.Do("SUNION", set1, set2))
	res := ""
	isFirst := true

	for _, itemV := range setMembers {
		tmpStr := fmt.Sprintf("%s", itemV)
		if isFirst {
			res += tmpStr
			isFirst = false
		} else {
			res += "|" + tmpStr
		}

	}
	return res
}

func GetAllGroupMembers(userName string, topic string) string {
	tmptopics := SgetAll(topic)
	topics := strings.Split(tmptopics, "|")
	numOfTopics := len(topics)
	var i int
	for i = 0; i < numOfTopics; i++ {
		isExist := SisExist(topics[i], userName)
		if isExist {
			break
		}
	}
	if i < numOfTopics {
		return SgetAll(topics[i])
	} else {
		return ""
	}
}

func HdelValue(key string) {
	conn.Do("DEL", key)
}

func DelUserInfo(userName string) {

	HdelValue(userName) //删除用户登录信息

	tmptopics := SgetAll("topic")
	topics := strings.Split(tmptopics, "|")

	numOfTopics := len(topics)

	var i int
	for i = 0; i < numOfTopics; i++ {
		isExist := SisExist(topics[i], userName)
		if isExist {
			break
		}
	}
	if i < numOfTopics {
		Srem(topics[i], userName)
	}
}
func SNumOfSet(key string) int {
	numsOfElements, err := redis.Int(conn.Do("SCARD", key))
	if err != nil {
		panic(err)
	}
	return numsOfElements
}

func SgetAllTopics(key string) string {
	setMembers, err := redis.Values(conn.Do("SMEMBERS", key))
	if err != nil {
		panic(err)
	}
	res := ""
	isFirst := true
	for _, itemV := range setMembers {
		tmpStr := fmt.Sprintf("%s", itemV)
		numsOfElements := SNumOfSet(tmpStr)
		tmpStr += ":"
		tmpStr += fmt.Sprintf("%d", numsOfElements)

		if isFirst {
			res += tmpStr
			isFirst = false
		} else {
			res += "|" + tmpStr
		}

	}
	return res
}

//服务端关闭的时候清空redis
func FlushRedis() {
	_, err := conn.Do("FLUSHDB")
	if err != nil {
		panic(err)
	}
}
