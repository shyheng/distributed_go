package dubbo

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/config"
	hessian "github.com/apache/dubbo-go-hessian2"
)

var (
	DayClient = &IDayService{} // 客户端指针
)

type IDayService struct {
	SelByDay func(ctx context.Context, s string) (Day, error)
	InsADD   func(ctx context.Context, day *Day) (bool, error)
}

func init() {
	hessian.RegisterPOJO(&Day{})
	config.SetConsumerService(DayClient)
}
