package service

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	spi "github.com/stigwang-gang/biz-demo/gomall/demo/demo_thrift/kitex_gen/spi"
)

type EchoService struct {
	ctx context.Context
} // NewEchoService new EchoService
func NewEchoService(ctx context.Context) *EchoService {
	return &EchoService{ctx: ctx}
}

// Run create note info
func (s *EchoService) Run(req *spi.Request) (resp *spi.Response, err error) {
	// Finish your business logic.
	info := rpcinfo.GetRPCInfo(s.ctx)
	fmt.Println(info.From().ServiceName())
	return &spi.Response{Message: req.Message}, nil
}
