package vo

import (
	"os"
	"strconv"
	"time"
	"util"
)

var (
	// web Server的Ip地址，默认为localhost
	Ip string
	// web Server的Ip地址，默认为9090
	Port string
	//  商品标识前缀
	Product_Pre string

	//商品1
	//  要卖的商品数量
	Product1_Max_Num int
	//  第一种商品productId
	Product1_Query_Name string
	//  第一种商品key 数据库标识
	Product1_Query_String string
	//  第一种商品开始秒杀的时间 格式为yyyy-MM-dd HH:mm:ss
	Product1_Begin_Time time.Time
	//  商品是否卖完的标识，优化使用
	Product1_Flag    bool
	Product1_Channel chan int
	Product1_Rel     map[int]int

	//商品2
	Product2_Max_Num      int
	Product2_Query_Name   string
	Product2_Query_String string
	Product2_Begin_Time   time.Time
	Product2_Flag         bool
	Product2_Channel      chan int
	Product2_Rel          map[int]int

	//商品3
	Product3_Max_Num      int
	Product3_Query_Name   string
	Product3_Query_String string
	Product3_Begin_Time   time.Time
	Product3_Flag         bool
	Product3_Channel      chan int
	Product3_Rel          map[int]int
)

func init() {
	//读取properties 文件中的各种配置参数
	timeLayout := "2006-01-02 15:04:05"  //转化所需模板
	loc, _ := time.LoadLocation("Local") //重要：获取时区
	myConfig := new(util.Config)
	path, _ := os.Getwd()
	path += "/project/conf/properties"
	myConfig.InitConfig(path)

	Ip = myConfig.Read("common", "ip")
	Port = myConfig.Read("common", "port")
	Product_Pre = myConfig.Read("common", "query_prefix")

	Product1_Max_Num, _ = strconv.Atoi(myConfig.Read("product_1", "max_num"))
	Product1_Query_Name = myConfig.Read("product_1", "query_name")
	Product1_Query_String = myConfig.Read("product_1", "total_query_name")
	Product1_Begin_Time, _ = time.ParseInLocation(timeLayout, myConfig.Read("product_1", "begin_time"), loc)

	Product2_Max_Num, _ = strconv.Atoi(myConfig.Read("product_2", "max_num"))
	Product2_Query_Name = myConfig.Read("product_2", "query_name")
	Product2_Query_String = myConfig.Read("product_2", "total_query_name")
	Product2_Begin_Time, _ = time.ParseInLocation(timeLayout, myConfig.Read("product_2", "begin_time"), loc)

	Product3_Max_Num, _ = strconv.Atoi(myConfig.Read("product_3", "max_num"))
	Product3_Query_Name = myConfig.Read("product_3", "query_name")
	Product3_Query_String = myConfig.Read("product_3", "total_query_name")
	Product3_Begin_Time, _ = time.ParseInLocation(timeLayout, myConfig.Read("product_3", "begin_time"), loc)

	Product1_Flag = true
	Product1_Channel = make(chan int, 200)
	Product1_Rel = make(map[int]int)

	Product2_Flag = true
	Product2_Channel = make(chan int, 200)
	Product2_Rel = make(map[int]int)

	Product3_Flag = true
	Product3_Channel = make(chan int, 200)
	Product3_Rel = make(map[int]int)
}
