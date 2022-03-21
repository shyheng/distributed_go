package main

import (
	"context"
	_ "dubbo.apache.org/dubbo-go/v3/common/logger" // dubbogo 框架日志
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports" // dubbogo 框架依赖，所有dubbogo进程都需要隐式引入一次
	"dubbo3-demo/dubbo"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//web.ShyRoute(r)

	r.GET("/json", func(c *gin.Context) {
		config.Load()
		userL, _ := dubbo.UserProviderClient.USER_LIST(context.TODO())
		//var da []dubbo.User
		//for i := 0; i < 10; i++ {
		//	da = append(da,dubbo.User{ID: int32(i), Name: "shy", Pass: "123", Del: 1})
		//}
		//data := dubbo.User{ID: 1, Name: "shy", Pass: "123", Del: 1}
		c.JSON(200, userL)

	})
	r.GET("/m", func(c *gin.Context) {
		config.Load()
		us, _ := dubbo.UserProviderClient.Us(context.TODO(), dubbo.User{12, "shy1", "cao", 1})
		c.JSON(200, us)
	})
	r.Run(":8050")
}
