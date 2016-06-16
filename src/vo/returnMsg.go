package vo

/**
 *实体类，封装秒杀结果的数据结构(对应seckilling接口)的返回值
 */
type ReturnMsg struct {
	/* 请求错误码  0 无错， 1有错 */
	Errno int `json:"errno"`
	/* 错误码，提醒信息   */
	ErrMsg string `json:"errMsg"`
}

func (msg *ReturnMsg) SetErrno(no int) {
	msg.Errno = no
}

func (msg *ReturnMsg) GetErrno() (res int) {
	return msg.Errno
}

func (msg *ReturnMsg) SetErrMsg(errMsg string) {
	msg.ErrMsg = errMsg
}

func (msg *ReturnMsg) GetErrMsg() (res string) {
	return msg.ErrMsg
}
