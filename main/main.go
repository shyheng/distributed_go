package main

import (
	"context"
	_ "dubbo.apache.org/dubbo-go/v3/common/logger" // dubbogo 框架日志
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports" // dubbogo 框架依赖，所有dubbogo进程都需要隐式引入一次
	"dubbo3-demo/dubbo"
	"dubbo3-demo/vue"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(vue.Cors())
	config.Load()
	//web.ShyRoute(r)

	r.GET("/json", func(c *gin.Context) {
		byDay, _ := dubbo.DayClient.SelByDay(context.TODO(), c.Query("day"))
		c.JSON(200, byDay)
	})
	r.GET("/cu", func(c *gin.Context) {
		s, _ := dubbo.DayClient.InsADD(context.TODO(), &dubbo.Day{
			c.Query("day"),
			c.Query("msg1"),
			c.Query("msg2"),
			c.Query("msg3"),
			c.Query("msg4"),
			c.Query("msg5"),
			c.Query("msg6"),
		})
		c.JSON(200, s)
	})
	r.Run(":8050")
}
