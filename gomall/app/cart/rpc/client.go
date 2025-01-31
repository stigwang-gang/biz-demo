package rpc

import (
	//"github.com/cloudwego/hertz/pkg/common/hlog"
	client "github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/stigwang-gang/biz-demo/gomall/app/cart/conf"
	cartUtils "github.com/stigwang-gang/biz-demo/gomall/app/cart/utils"
	"github.com/stigwang-gang/biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"sync"
)

var (
	ProductClient productcatalogservice.Client
	once          sync.Once
)

func Init() {
	once.Do(func() {
		initProductClient()
	})
}

func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	cartUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	cartUtils.MustHandleError(err)
}
