package dubbo

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/config"
)

var (
	UserSerC = &UserSer{} // 客户端指针
)

type UserSer struct {
	S func(ctx context.Context, id int32) (string, error)
}

func init() {
	config.SetConsumerService(UserSerC)
}
