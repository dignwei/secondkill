package main

import (
	"controller"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"vo"
)

func main() {

	go listen(vo.Product1_Rel, vo.Product1_Channel, &vo.Product1_Flag, vo.Product1_Max_Num)
	go listen(vo.Product2_Rel, vo.Product2_Channel, &vo.Product2_Flag, vo.Product2_Max_Num)
	go listen(vo.Product3_Rel, vo.Product3_Channel, &vo.Product3_Flag, vo.Product3_Max_Num)

	runtime.GOMAXPROCS(2)
	http.HandleFunc("/zaixianshang/queryUserSeckillingInfo", controller.QueryUserSeckillingInfo)       //设置访问的路由
	http.HandleFunc("/zaixianshang/seckilling", controller.Seckilling)                                 //设置访问的路由
	http.HandleFunc("/zaixianshang/queryProductSeckillingInfo", controller.QueryProductSeckillingInfo) //设置访问的路由
	err := http.ListenAndServe(":9090", nil)                                                           //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func listen(rel map[int]int, c chan int, flag *bool, num int) {
	fmt.Println("init listen!")
	fmt.Println("商品总数: ", num)
	for i := 0; i < num; i++ {
		x := <-c
		rel[x] = i
		//fmt.Println("在listen for 循环里")
	}
	//fmt.Println("跳出循环,应该将flag 变成false")
	*flag = false
	close(c)
}
