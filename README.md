##DMall
####团队介绍
+ 队名：DMall
+ 队号：箭在弦上，一击必中

####产品介绍
DMall是一个基于Go语言的商品秒杀工具

功能：

+ 提供统一的HTTP接口：秒杀接口、秒杀结果接口（个人及全部）
+ 能正确运行程序，输出正确的秒杀结果（100件商品被正确秒完）
+ 能正确运行程序，输出正确的秒杀结果（100件商品被正确秒完）

性能

+ 系统响应无明显延迟（100ms内），不会发生接口无响应、卡死等情况

+ 支持多并发请求数，服务器稳定


文档：

	设计文档

	项目安排


使用说明

+ Linux安装命令如下：

	wget http://download.redis.io/redis-stable.tar.gz

	tar xzf redis-stable.tar.gz

	cd redis-stable

	make

+ Windows安装较为简单：

	下载安装后到redis安装目录下运行redis-server.exe 即可，redis默认运行端口为：6379

DMall Linux下安装

+ 进入目录启动load.sh脚本

+ 网页直接打开http://localhost:9090/secondkill/seckilling?productid=111&userid=1接口进行秒杀
+  DMall提供了iOS和Android客户端使用
