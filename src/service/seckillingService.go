package service

import (
	"fmt"
	"strconv"
	"time"
	"vo"
)

func ServiceSeckilling(userid, productid string) *vo.ReturnMsg {
	message := &vo.ReturnMsg{0, ""} //返回消息

	var channel *chan int
	var flag *bool
	var rel map[int]int
	var beginTime time.Time

	//选择商品种类
	switch productid {
	case vo.Product1_Query_Name:
		flag = &vo.Product1_Flag
		channel = &vo.Product1_Channel
		rel = vo.Product1_Rel
		beginTime = vo.Product1_Begin_Time
	case vo.Product2_Query_Name:
		flag = &vo.Product2_Flag
		channel = &vo.Product2_Channel
		rel = vo.Product2_Rel
		beginTime = vo.Product2_Begin_Time
	case vo.Product3_Query_Name:
		flag = &vo.Product3_Flag
		channel = &vo.Product3_Channel
		rel = vo.Product3_Rel
		beginTime = vo.Product3_Begin_Time
	default:
		message.SetErrno(1)
		message.SetErrMsg("所要秒杀的商品不存在，秒杀失败")
		return message
	}

	//增加时间控制
	if time.Now().Before(beginTime) {
		message.SetErrno(1)
		message.SetErrMsg("秒杀尚未开始，请稍后再来")
		return message
	}

	fmt.Println("*flag = ", *flag)

	//商品已经卖光,秒杀直接返回失败
	if !*flag {
		message.SetErrno(0)
		message.SetErrMsg("秒杀失败")
		return message
	}

	uId, err := strconv.Atoi(userid)
	if err != nil {
		message.SetErrno(1)
		message.SetErrMsg("用户信息异常，秒杀失败")
		return message
	}

	//防重入
	if _, ok := rel[uId]; ok {
		message.SetErrno(0)
		message.SetErrMsg("已经秒杀过该产品,请稍后查询秒杀结果~")
		return message
	}

	//写入当前订单的channel
	*channel <- uId //channel关闭后，此步骤会抛出panic
	message.SetErrno(0)
	message.SetErrMsg("秒杀中，请稍后查询")

	/* //测试商品映射关系
	fmt.Println("start--------当前map长度:", len(rel))
	for Key, Value := range vo.Product1_Rel {
		fmt.Println("key = ", Key, "value = ", Value)
	}
	fmt.Println("end---------当前map长度:", len(rel))
	*/

	return message
}
