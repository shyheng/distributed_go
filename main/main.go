package main

import (
	_ "dubbo.apache.org/dubbo-go/v3/common/logger" // dubbogo 框架日志
	_ "dubbo.apache.org/dubbo-go/v3/imports"       // dubbogo 框架依赖，所有dubbogo进程都需要隐式引入一次
	"dubbo3-demo/web"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	web.ShyRoute(r)
	r.Run(":8050")
}
