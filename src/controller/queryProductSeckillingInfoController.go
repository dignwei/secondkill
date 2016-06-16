package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"service"
	"strings"
	"vo"
)

//返回商品秒杀结果，首先解析请求参数，再调用ServiceQueryProductSeckillingInfo获得商品的秒杀结果
func QueryProductSeckillingInfo(resp http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	counter += 1
	fmt.Println("Server 已收到的请求总数 : ", counter)
	message := &vo.ResultProductMsg{0, "", nil}
	defer func() { //异常处理
		if err := recover(); err != nil {
			message.SetErrno(1)
			message.SetErrMsg("内部错误")
			message.SetList(nil)
		}
		if jsonstr, jsonerr := json.Marshal(message); jsonerr == nil {
			fmt.Fprintf(resp, string(jsonstr))
		} else {
			fmt.Fprintf(resp, "json错误")
		}
	}()
	productid := ""
	for key, value := range req.Form {
		if key == "productid" {
			productid = strings.Join(value, "")
		}
	}
	message = service.ServiceQueryProductSeckillingInfo(productid)
}
