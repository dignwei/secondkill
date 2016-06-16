package test

import (
	"service"
	"testing"
	"vo"
	"fmt"
)
// userid 和 productid 参数错误 case
func Test_ServiceSeckilling_1(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log("用例1测试通过！")
		} else {
			t.Error("用例1测试未通过！")
		}
	}()
	service.ServiceSeckilling("", "")
}

// 秒杀失败 case
func Test_ServiceSeckilling_2(t *testing.T) {
	vo.Flag = false
	message := service.ServiceSeckilling("123", "111")
	if message.GetErrMsg() == "秒杀失败" && message.GetErrno() == 0 {
		t.Log("用例2测试通过！")
	} else {
		t.Error("用例2测试未通过！")
	}
}

// productid不存在 case
func Test_ServiceSeckilling_3(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log("用例3测试通过！")
		} else {
			t.Error("用例3测试未通过！")
		}
	}()
	service.ServiceSeckilling("123", "123")
}

// 秒杀中 case
func Test_ServiceSeckilling_4(t *testing.T) {
	message := &vo.ReturnMsg{0, ""}
	defer func() {
		fmt.Println(message.GetErrMsg())
		if err := recover(); err != nil {
			t.Error("用例4测试未通过！")
		} else {
			t.Log("用例4测试通过！")
		}
	}()
	message = service.ServiceSeckilling("123", "111")
}