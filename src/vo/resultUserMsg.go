package vo

/**
 *实体类，封装用户查询自己秒杀结果的数据结构(对应queryUserSeckillingInfo接口)的返回值
 */

type ResultPersonMsg struct {
	/* 请求错误码  0 无错， 1有错 */
	Errno int `json:"errno"`
	/*  查询接口的状态码 0 秒杀还未开始，1 秒杀成功，2 秒杀失败，3 秒杀中 4 非法信息，无资格秒杀*/
	Status string `json:"status"`
	/*  请求码为0 & 状态码为1 时，标识该用户购买商品编号,否则，存储错误提醒信息*/
	GoodsId string `json:"goodsid"`

	ProductId string `json:"productid"`
}

func (msg *ResultPersonMsg) SetErrno(no int) {
	msg.Errno = no
}

func (msg *ResultPersonMsg) GetErrno() (res int) {
	return msg.Errno
}

func (msg *ResultPersonMsg) SetStatus(status string) {
	msg.Status = status
}

func (msg *ResultPersonMsg) GetStatus() (status string) {
	return msg.Status
}

func (msg *ResultPersonMsg) SetGoodsId(goodsId string) {
	msg.GoodsId = goodsId
}

func (msg *ResultPersonMsg) GetGoodsId() (goodsId string) {
	return msg.GoodsId
}

func (msg *ResultPersonMsg) SetProductId(productId string) {
	msg.ProductId = productId
}

func (msg *ResultPersonMsg) GetProductId() (productId string) {
	return msg.ProductId
}
