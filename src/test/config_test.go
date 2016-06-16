package test

import (
	"os"
	"util"
	"testing"
)

func Test_Config(t *testing.T) {
	myConfig := new(util.Config)
	path, _ := os.Getwd()
	myConfig.InitConfig(path)
	if path == "E:\\IdeaProjects\\zaixianshang\\project" {
		t.Log("配置测试通过")
	} else {
		t.Error("配置测试未通过")
	}
}
