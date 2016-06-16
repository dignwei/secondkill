package test
import (
	"testing"
	"service"
	"vo"
	"fmt"
)

// userid 和 productid 参数错误 case
func TestQueryUserSeckillingInfo_1(t *testing.T){
	userid,productid :="",""
	defer func() {
		if err := recover(); err != nil {
			t.Log("Test_QueryUserSeckillingInfo_1 参数校验 测试通过")
		}
	}()
	service.QueryUserSeckillingInfo(userid,productid)
	t.Error("Test_QueryUserSeckillingInfo_1 参数校验 测试失败") 
}
 // userid正常,productID 不是第一种商品id(vo.Product2_Query_Name）
func TestQueryUserSeckillingInfo_2(t *testing.T){

	userid,productid :="userid","222"
	defer func() {
		if err := recover(); err != nil {
			t.Log("Test_QueryUserSeckillingInfo_2 只卖 "+vo.Product1_Query_Name+" 商品，测试通过")
		}
	}()
	service.QueryUserSeckillingInfo(userid,productid)
	t.Error("Test_QueryUserSeckillingInfo_1 只卖 "+vo.Product1_Query_Name+" 商品，测试失败") 
}

// userid正常，秒杀成功，依赖 redis 中有userid=1,已经秒中了productid=111,redis 通过 conf/properties配置
func TestQueryUserSeckillingInfo_3(t *testing.T){
	userid,productid :="1","111"
	message := &vo.ResultPersonMsg{0, "", ""}
	message = service.QueryUserSeckillingInfo(userid,productid)
	fmt.Println( message.GetStatus());
	fmt.Println( message.GetErrno());
	if 0 == message.GetErrno() && "1" == message.GetStatus(){
		t.Log("商品秒杀成功 测试通过")
	} else {
		t.Error("商品秒杀成功 测试失败")
	}
}

// userid正常，秒杀中，依赖 redis  不存在userid=2且productid=111的商品还没卖完,redis 通过 conf/properties
func TestQueryUserSeckillingInfo_4(t *testing.T){

	userid,productid :="2","111"
	message := &vo.ResultPersonMsg{0, "", ""}
	message = service.QueryUserSeckillingInfo (userid,productid)
	if 0 == message.GetErrno() && "3" == message.GetStatus(){
		t.Log("商品秒杀中 测试通过")
	} else {
		t.Error("商品秒杀中 测试失败")
	}
}

// userid正常，秒杀未开始，依赖 redis  不存在userid=2且productid=111的商品还没卖完,redis 通过 conf/properties
func TestQueryUserSeckillingInfo_6(t *testing.T){

	userid,productid :="2","111"
	message := &vo.ResultPersonMsg{0, "", ""}
	message = service.QueryUserSeckillingInfo (userid,productid)
	if 0 == message.GetErrno() && "0" == message.GetStatus(){
		t.Log("商品秒杀未开始 测试通过")
	} else {
		t.Error("商品未开始 测试失败")
	}

}


// userid正常，秒杀失败，依赖 redis  不存在userid=1000，且 productid=111的商品已经卖完,redis 通过 conf/properties
func TestQueryUserSeckillingInfo_5(t *testing.T){

	userid,productid :="1000","111"
	message := &vo.ResultPersonMsg{0, "", ""}
	message = service.QueryUserSeckillingInfo (userid,productid)
	if 0 == message.GetErrno() && "2" == message.GetStatus(){
		t.Log("商品秒杀失败 测试通过")
	} else {
		t.Error("商品秒杀失败 测试失败")
	}
}
