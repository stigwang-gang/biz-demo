package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	//common "github.com/stigwang-gang/biz-demo/gomall/app/frontend/hertz_gen/frontend/common"
	product "github.com/stigwang-gang/biz-demo/gomall/app/frontend/hertz_gen/frontend/product"
	"github.com/stigwang-gang/biz-demo/gomall/app/frontend/infra/rpc"
	rpcproduct "github.com/stigwang-gang/biz-demo/gomall/rpc_gen/kitex_gen/product"
)

type GetProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductService(Context context.Context, RequestContext *app.RequestContext) *GetProductService {
	return &GetProductService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductService) Run(req *product.ProductReq) (resp map[string]any, err error) {
	p, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"item": p.Product,
	}, nil
}
