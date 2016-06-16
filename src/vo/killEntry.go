package vo

/**
 * 实体类，映射redis中的<商品编号，用户标识>关系
 *
 */
type KillEntry struct {
	/* 用户标识 */
	Userid string `json:"userid"`
	/* 商品标识 */
	Goodsid string `json:"goodsid"`
}

func (entry *KillEntry) SetUserid(userid string) {
	entry.Userid = userid
}

func (entry *KillEntry) SetGoodsid(goodsid string) {
	entry.Goodsid = goodsid
}

func (entry *KillEntry) GetGoodsid() (res string) {
	return entry.Goodsid
}

func (entry *KillEntry) GetUserid() (res string) {
	return entry.Userid
}
