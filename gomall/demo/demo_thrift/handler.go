package main

import (
	"context"
	"github.com/stigwang-gang/biz-demo/gomall/demo/demo_thrift/biz/service"
	spi "github.com/stigwang-gang/biz-demo/gomall/demo/demo_thrift/kitex_gen/spi"
)

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// Echo implements the EchoImpl interface.
func (s *EchoImpl) Echo(ctx context.Context, req *spi.Request) (resp *spi.Response, err error) {
	resp, err = service.NewEchoService(ctx).Run(req)

	return resp, err
}
