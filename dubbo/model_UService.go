package dubbo

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/config"
)

var (
	UserClient = &UService{} // 客户端指针
)

type UService struct {
	STRING func(ctx context.Context, id int32) (string, error)
}

func init() {
	config.SetConsumerService(UserClient)
}
