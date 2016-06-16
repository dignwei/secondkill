package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"service"
	"strings"
	"vo"
)

var counter = 0 //记录请求数，非线程安全，并发量较大时不保证精确

func Seckilling(resp http.ResponseWriter, req *http.Request) {
	req.ParseForm() //解析参数，默认是不会解析的
	counter += 1
	fmt.Println("第 ", counter, " 个数据")
	message := &vo.ReturnMsg{0, ""}
	defer func() { //异常处理
		if err := recover(); err != nil {
			message.SetErrno(1)
			message.SetErrMsg("很遗憾，秒杀失败了") //目前底层只有在channel关闭时才会出现此panic
		}
		if jsonstr, jsonerr := json.Marshal(message); jsonerr == nil {
			fmt.Fprintf(resp, string(jsonstr))
		} else {
			fmt.Fprintf(resp, "json错误")
		}
	}()
	userid, productid := "", ""
	for key, value := range req.Form {
		if key == "userid" {
			userid = strings.Join(value, "")
		} else if key == "productid" {
			productid = strings.Join(value, "")
		}
		fmt.Print("key:", key)
		fmt.Println("  value:", strings.Join(value, ""))
	}

	message = service.ServiceSeckilling(userid, productid)
}
