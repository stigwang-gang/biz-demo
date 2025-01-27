package service

import (
	"context"
	spi "github.com/stigwang-gang/biz-demo/gomall/demo/demo_thrift/kitex_gen/spi"
	"testing"
)

func TestEcho_Run(t *testing.T) {
	ctx := context.Background()
	s := NewEchoService(ctx)
	// init req and assert value

	req := &spi.Request{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
