package dubbo

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/config"
	hessian "github.com/apache/dubbo-go-hessian2"
)

var (
	UserProviderClient = &UserService{}
)

type UserService struct {
	Test     func(ctx context.Context, id int32) (User, error)
	USERS    func(ctx context.Context) (Users, error)
	UserList func(ctx context.Context) (User, error)
}

func init() {
	hessian.RegisterPOJO(&Users{})
	hessian.RegisterPOJO(&User{})
	config.SetConsumerService(UserProviderClient)
}
