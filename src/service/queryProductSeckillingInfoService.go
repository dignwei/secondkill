package service

import (
	"fmt"
	"strconv"
	"time"
	"vo"
)

//根据商品id查询成功秒杀的所有用户id和用户购买商品的具体编号
func ServiceQueryProductSeckillingInfo(productid string) *vo.ResultProductMsg {
	returnMsg := &vo.ResultProductMsg{0, "", nil}

	var rel map[int]int
	var maxNum int //添加秒杀商品数量作为切片容量，优化效率
	var beginTime time.Time

	switch productid {
	case vo.Product1_Query_Name:
		rel = vo.Product1_Rel
		maxNum = vo.Product1_Max_Num
		beginTime = vo.Product1_Begin_Time
	case vo.Product2_Query_Name:
		rel = vo.Product2_Rel
		maxNum = vo.Product2_Max_Num
		beginTime = vo.Product2_Begin_Time
	case vo.Product3_Query_Name:
		rel = vo.Product3_Rel
		maxNum = vo.Product3_Max_Num
		beginTime = vo.Product3_Begin_Time
	default:
		fmt.Println("errMsg:", "productid不存在")
		returnMsg.SetErrno(1)
		returnMsg.SetErrMsg("productid不存在")
		returnMsg.SetList(nil)
		return returnMsg
	}

	if time.Now().Before(beginTime) {
		fmt.Println("errMsg:", "秒杀尚未开始，无法查询")
		returnMsg.SetErrno(1)
		returnMsg.SetErrMsg("秒杀尚未开始，无法查询")
		returnMsg.SetList(nil)
		return returnMsg
	}

	if len(rel) == 0 {
		returnMsg.SetErrno(1)
		returnMsg.SetErrMsg("秒杀尚未开始或是内部错误") //秒杀系统不可能卖不出去商品
		returnMsg.SetList(nil)
		return returnMsg
	}

	goodsList := make([]vo.KillEntry, 0, maxNum)
	for key, value := range rel {
		userid := strconv.Itoa(key)
		goodsid := strconv.Itoa(value)
		killEntry := vo.KillEntry{userid, goodsid}
		goodsList = append(goodsList, killEntry)
	}

	returnMsg.SetErrno(0)
	returnMsg.SetList(goodsList)
	return returnMsg

	/*
		//原redis方式获取数据
		productInfo, _ := dao.RedisPoolOne.LRange(vo.Product1_Query_String)

		if productInfo == nil {
			fmt.Println("errMsg:", "无法查询到结果")
			panic("无法查询到结果")
		}
		goodsList := []vo.KillEntry{}
		for _, entry := range productInfo {
			tmp := strings.Split(entry, "*")
			userid := tmp[0]
			goodsid := tmp[1]
			killEntry := vo.KillEntry{userid, goodsid}
			goodsList = append(goodsList, killEntry)
		}
		returnMsg.SetErrno(0)
		returnMsg.SetList(goodsList)
		return returnMsg
	*/
}
