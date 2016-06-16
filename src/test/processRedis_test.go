package test

import(
	"testing"
	"dao"
)

func Test_ProcessRedis_1(t *testing.T){//测试线程池的入队和出队
	dao.OpenRedis("192.168.2.186","6379")
	dao.LPushValue("111","1")
	if res := dao.LPopValue("111"); res != "1"{
		t.Error("测试1:redis入队出队测试失败!")
	} else{
		t.Log("测试1:redis入队出队测试通过!")
	}

}

func Test_ProcessRedis_2(t *testing.T){//测试线程池的出队顺序
	dao.OpenRedis("192.168.2.186","6379")
	dao.LPushValue("111","1")
	dao.LPushValue("222","2")
	if res := dao.LPopValue("222"); res != "2"{
		t.Error("测试2:线程池出队顺序测试失败!")
	} else {
		t.Log("测试2:线程池出队顺序测试通过!")
	}
}