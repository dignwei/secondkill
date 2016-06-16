package test

import (
	"dao"
	"testing"
	"fmt"
)

func Test_Redis(t *testing.T) {

	redisone := dao.RedisPoolOne
	//redis的数据输入输出
	redisone.Set("123", "2123")
	resultone, _ :=redisone.Get("123")
	fmt.Println(resultone)

	//redis list数据输入输出
	redisone.RPush("1243", "1")
	redisone.RPush("1243", "2")

	resulttwo, _:=redisone.RPop("1243")
	resultthree, _:=redisone.RPop("1243")
	if resultone=="2123"{
		t.Log("测试一通过")
	}else{
		t.Error(resultone)
	}
	if resulttwo=="2"{
		t.Log("测试二通过")
	}else{
		t.Error(resulttwo)
	}
	if resultthree=="1"{
		t.Log("测试三通过")
	}else{
		t.Error(resultthree)
	}
}
