package main

import (
	"fmt"
	"time"
)

//使用redis队列实现的消费功能
/*
func main() {
	defer dao.CloseRedis()
	defer func() {
		// 必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("listen the world!")

	dao.OpenRedis(vo.Ip, vo.Port)
	//初始化商品数量
	dao.HSetValue(vo.Product_Pre+vo.Product1_Query_Name, "0")

	for {
		//超过商品总数,停止监听
		//todo 添加配置文件,尝试多商品
		if count, err := strconv.Atoi(dao.HGetValue(vo.Product_Pre + vo.Product1_Query_Name)); count >= vo.Product1_Max_Num {
			if err != nil {
				panic(err)
			}
			fmt.Println("listen finish!")
			break
		}
		popValue := dao.LPopValue("list") //消费队列
		if popValue != "" {
			fmt.Print("popValue:" + popValue)
			var qe vo.QueueEntry
			var userId string
			//解压数据
			if err := json.Unmarshal([]byte(popValue), &qe); err == nil {
				fmt.Println(qe.Userid)
				userId = qe.Userid
			} else {
				continue //解压失败则淘汰该用户,这里不抛出异常
			}
			if checkValid(userId) {
				incCount(vo.Product_Pre+vo.Product1_Query_Name, userId) //写入redis
			}
			//fmt.Print(popValue + "\n")
		}
	}
}


func checkValid(userId string) bool {
	if dao.HGetValue(userId) != "" {
		return false
	}
	return true
}

//商品计数器加1,并存储,异常均在异步方法中打印
func incCount(productId string, userId string) {
	//商品数量计数器加一
	goodsid := 1
	tmp := dao.HGetValue(productId)
	//fmt.Print("  "+tmp)
	//fmt.Print("\n")
	b, _ := strconv.Atoi(tmp)
	goodsid = b + 1
	dao.HSetValue(productId, strconv.Itoa(goodsid))

	//存储用户和其购买的商品关系
	dao.HSetValue(userId, strconv.Itoa(goodsid))
	fmt.Print(userId + "****" + strconv.Itoa(goodsid))
	fmt.Print("\n")
	newValue := userId + "*" + strconv.Itoa(goodsid)
	fmt.Print(newValue)
	fmt.Print("\n")
	//向全部购买信息中push数据
	dao.LPushValue(vo.Product1_Query_String, newValue)
}
*/
func main() {

	//var i = false
	//func(j *bool) {
	//	*j = true
	//}(&i)
	//
	//fmt.Println(i)
	//
	//var x = 100
	//var y *int
	//y = &x
	//*y = 1
	//fmt.Println(x)
	//fmt.Println(*y)

	toBeCharge := "2016-06-14 11:02:15"                             //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	timeLayout := "2006-01-02 15:04:05"                             //转化所需模板
	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()                                            //转化为时间戳 类型是int64
	fmt.Println(theTime)                                            //打印输出theTime 2015-01-01 15:15:00 +0800 CST
	fmt.Println(sr)
	fmt.Println(time.Now().Unix())
	fmt.Println(theTime.Before(time.Now()))
}
