package main

import (
	_ "dubbo.apache.org/dubbo-go/v3/common/logger" // dubbogo 框架日志
	_ "dubbo.apache.org/dubbo-go/v3/imports"       // dubbogo 框架依赖，所有dubbogo进程都需要隐式引入一次
	"dubbo3-demo/vue"
	"github.com/gin-gonic/gin"
)

type Time struct {
	Day string
	Msg string
}

func main() {

	r := gin.Default()
	r.Use(vue.Cors())
	//web.ShyRoute(r)

	//r.GET("/json", func(c *gin.Context) {
	//	config.Load()
	//	userL, _ := dubbo.UserProviderClient.USER_LIST(context.TODO())
	//	//var da []dubbo.User
	//	//for i := 0; i < 10; i++ {
	//	//	da = append(da,dubbo.User{ID: int32(i), Name: "shy", Pass: "123", Del: 1})
	//	//}
	//	//data := dubbo.User{ID: 1, Name: "shy", Pass: "123", Del: 1}
	//	c.JSON(200, userL)
	//
	//})
	//r.GET("/m", func(c *gin.Context) {
	//	config.Load()
	//	us, _ := dubbo.UserProviderClient.Us(context.TODO(), dubbo.User{12, "shy1", "cao", 1})
	//	c.JSON(200, us)
	//})

	r.GET("/json", func(context *gin.Context) {
		context.JSON(200, Time{"2022-03-21", "睡觉了"})
	})

	r.Run(":8050")
}
