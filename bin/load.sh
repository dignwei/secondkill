#!/bin/bash
sudo ulimit -n 10032 -u 1024
echo "Init..."
ps -ef | grep ./webServer | grep -v grep | awk '{print $2}' | xargs kill -9
ps -ef | grep ./listen | grep -v grep | awk '{print $2}' | xargs kill -9
ps -ef | grep redis-server | grep -v grep | awk '{print $2}' | xargs kill -9
echo "init done!"
echo "service start..."
#export GOPATH=/Users/didi/trainningProject/zaixianshang
echo "start redis-server..."
redis-server &
cp -r ../conf  ./ 
echo "start webServer..."
go build ../src/main/webServer.go 
./webServer &
echo "start listener..." 
go build ../src/main/listen.go
./listen &
echo "done!"

