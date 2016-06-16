package test

/**测试每次都要对redis服务器清空
*商品查询接口单元测试
*提供覆盖三种情况：productid错误；productid正确单还没有开始秒杀；productid正确且已经开始秒杀
 */

import (
	"dao"
	"encoding/json"
	"fmt"
	"service"
	"testing"
	"vo"
)

func Test_ServiceQueryProductSeckillingInfo_1(t *testing.T) {

	one := productidIn("123")
	fmt.Println(one)
	two := productidIn("111")
	fmt.Println(two)
	dao.RedisPoolOne.RPush(vo.Product1_Query_String, "1*1")
	three := productidIn("111")
	if one == "" {
		t.Log("案例一通过")
	} else {
		t.Error(one)
	}
	if two == "{\"errno\":0,\"list\":[]}" {
		t.Log("案例二通过")
	} else {
		t.Error(two)
	}
	if three == "{\"errno\":0,\"list\":[{\"userid\":\"1\",\"goodsid\":\"1\"}]}" {
		t.Log("案例三通过")
	} else {
		t.Error(three)
	}
}
func productidIn(productid string) string {
	message := &vo.ResultProductMsg{0, nil}
	returnmes := "pwd"
	defer func() { //异常处理
		if err := recover(); err != nil {
			returnmes = "kong"
		}
		if _, jsonerr := json.Marshal(message); jsonerr == nil {
			returnmes = "json"
		} else {
			returnmes = "jsonwrong"
		}

	}()
	message = service.ServiceQueryProductSeckillingInfo(productid)
	returnme, _ := json.Marshal(message)
	returnmes = string(returnme)
	return returnmes

}
