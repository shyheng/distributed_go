package main

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/logger" // dubbogo 框架日志
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports" // dubbogo 框架依赖，所有dubbogo进程都需要隐式引入一次
	"dubbo3-demo/api"
)

// export DUBBO_GO_CONFIG_PATH=dubbogo.yml 运行前需要设置环境变量，指定配置文件位置
func main() {
	// 启动框架
	if err := config.Load(); err != nil {
		panic(err)
	}
	//var i int32 = 1
	// 发起调用
	//user, err := api.UserProviderClient.GetUser(context.TODO(), i)
	user, err := api.UserProviderClient.Test(context.TODO(), 2)
	if err != nil {
		panic(err)
	}
	logger.Infof("response result: %+v", user.ID)
}
