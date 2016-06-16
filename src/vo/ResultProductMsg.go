package vo

/**
 * 实体类，映射查询指定商品id的结果的数据结构(对应queryProductSeckillingInfo接口)的返回值
 *
 */
type ResultProductMsg struct {
	/* 请求错误码  0 无错， 1有错 */
	Errno int `json:"errno"`
	/* 错误信息 */
	ErrMsg string `json:"errMsg"`
	/* 商品标识 和 用户标识 二元组集合，如果为nil，表示现在无商品卖出 */
	List []KillEntry `json:"list"`
}

func (rpm *ResultProductMsg) SetErrMsg(errMsg string) {
	rpm.ErrMsg = errMsg
}

func (rpm *ResultProductMsg) GetErrMsg() (msg string) {
	return rpm.ErrMsg
}

func (rpm *ResultProductMsg) SetErrno(errno int) {
	rpm.Errno = errno
}

func (rpm *ResultProductMsg) SetList(list []KillEntry) {
	rpm.List = list
}

func (rpm *ResultProductMsg) GetErrno() (res int) {
	return rpm.Errno
}

func (rpm *ResultProductMsg) GetList() (list []KillEntry) {
	return rpm.List
}
