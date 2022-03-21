package web

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/config"
	"dubbo3-demo/dubbo"
	"github.com/gin-gonic/gin"
)

type Shy struct {
	Name string
}

func (s Shy) Add(ctx *gin.Context) {
	ctx.String(200, s.Name)
}

func ShyRoute(engine *gin.Engine) {

	routerGroup := engine.Group("/")
	{
		config.Load()
		users, _ := dubbo.UserProviderClient.USERS(context.TODO())

		user, _ := dubbo.UserProviderClient.Test(context.TODO(), 1)
		u, _ := dubbo.UserClient.STRING(context.TODO(), 1)
		//k, _ := dubbo.UserSerC.S(context.TODO(),1)
		routerGroup.GET("/add", Shy{Name: users.Name}.Add)
		routerGroup.GET("/ad", Shy{Name: user.Name}.Add)
		routerGroup.GET("/u", Shy{Name: u}.Add)
		//routerGroup.GET("/k",Shy{Name: userL[0:1]}.Add)
	}

}
