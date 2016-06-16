package service

import (
	"fmt"
	"strconv"
	"time"
	"vo"
)

/*
 *   用户查询自己的秒杀结果接口
 *
 * errno  : 1
 *        status :2 参数错误
 * errno  : 0
 *        status :0  秒杀还未开始 ,商品被卖了
 *        status :1  秒杀成功    , 成功秒杀到，redis中查询到goodsId
 *        status :2  秒杀失败    , 没有秒杀到，redis中未查询到goodsId，且商品已经被卖完
 *        status :3  在秒杀中    , 没有秒杀到，redis中未查询到goodsId，但是商品还未卖完
 *
 */
func QueryUserSeckillingInfo(userid, productid string) *vo.ResultPersonMsg {
	retMessage := &vo.ResultPersonMsg{0, "", "", ""}

	var beginTime time.Time

	//以下情况属于可预测失败，不应以异常方式解决，而应直接返回错误
	if userid == "" || productid == "" {
		fmt.Println("errMsg:", "参数错误")
		retMessage.SetErrno(1)
		retMessage.SetStatus("4")
		retMessage.SetGoodsId("缺少必要信息，秒杀失败")
		retMessage.SetProductId(productid)
		return retMessage
	}

	var rel map[int]int
	var flag *bool

	switch productid {
	case vo.Product1_Query_Name:
		flag = &vo.Product1_Flag
		rel = vo.Product1_Rel
		beginTime = vo.Product1_Begin_Time
	case vo.Product2_Query_Name:
		flag = &vo.Product2_Flag
		rel = vo.Product2_Rel
		beginTime = vo.Product2_Begin_Time
	case vo.Product3_Query_Name:
		flag = &vo.Product3_Flag
		rel = vo.Product3_Rel
		beginTime = vo.Product3_Begin_Time
	default:
		retMessage.SetErrno(1)
		retMessage.SetStatus("4")
		retMessage.SetGoodsId("所要秒杀的商品不存在，秒杀失败")
		retMessage.SetProductId(productid)
		return retMessage
	}

	//增加时间控制
	if time.Now().Before(beginTime) {
		retMessage.SetErrno(1)
		retMessage.SetStatus("4")
		retMessage.SetGoodsId("秒杀尚未开始，无法查询")
		retMessage.SetProductId(productid)
		return retMessage
	}

	uId, err := strconv.Atoi(userid)
	if err != nil {
		retMessage.SetErrno(1)
		retMessage.SetStatus("4")
		retMessage.SetGoodsId("用户信息异常，秒杀失败")
		retMessage.SetProductId(productid)
		return retMessage
	}

	fmt.Println("*flag = ", *flag)

	if key, ok := rel[uId]; ok {
		//查询到数据
		goodsid := strconv.Itoa(key)
		retMessage.SetErrno(0)
		retMessage.SetStatus("1")
		retMessage.SetGoodsId(goodsid)
		retMessage.SetProductId(productid)
	} else {
		//当前未命中
		if *flag {
			//未秒杀完毕,仍有机会
			retMessage.SetErrno(0)
			retMessage.SetStatus("3")
			retMessage.SetGoodsId("在秒杀中，请稍后查询...")
			retMessage.SetProductId(productid)
		} else {
			//秒杀失败
			retMessage.SetErrno(0)
			retMessage.SetStatus("2")
			retMessage.SetGoodsId("秒杀失败,未秒杀到商品")
			retMessage.SetProductId(productid)
		}
	}
	return retMessage
}
