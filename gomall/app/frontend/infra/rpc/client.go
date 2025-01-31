package rpc

import (
	//"github.com/cloudwego/hertz/pkg/common/hlog"
	client "github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/stigwang-gang/biz-demo/gomall/app/frontend/conf"
	frontendUtils "github.com/stigwang-gang/biz-demo/gomall/app/frontend/utils"
	"github.com/stigwang-gang/biz-demo/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/stigwang-gang/biz-demo/gomall/rpc_gen/kitex_gen/user/userservice"
	"sync"
)

var (
	UserClient    userservice.Client
	ProductClient productcatalogservice.Client
	once          sync.Once
)

func Init() {
	once.Do(func() {
		initUserClient()
		initProductClient()
	})
}
func initUserClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	frontendUtils.MustHandleError(err)
}

func initProductClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	frontendUtils.MustHandleError(err)
	opts = append(opts, client.WithResolver(r))
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	frontendUtils.MustHandleError(err)
}
