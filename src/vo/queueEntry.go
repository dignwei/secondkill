package vo

/**
 * 实体类，映射redis中list队列中的每一项
 *
 */
type QueueEntry struct {
	/* 用户标识 */
	Userid string `json:"userid"`
	/* 要秒杀商品标识 */
	Productid string `json:"productid"`
	/* 请求时间戳 */
	Time string `json:"time"`
}

func (queueEntry *QueueEntry) GetUserid() (str string) {
	return queueEntry.Userid
}
func (queueEntry *QueueEntry) GetProductid() (str string) {
	return queueEntry.Productid
}
func (queueEntry *QueueEntry) GetTime() (str string) {
	return queueEntry.Time
}

func (queueEntry *QueueEntry) SetUserid(userid string) {
	queueEntry.Userid = userid
}
func (queueEntry *QueueEntry) SetProductid(productid string) {
	queueEntry.Productid = productid
}
func (queueEntry *QueueEntry) SetTime(time string) {
	queueEntry.Time = time
}
