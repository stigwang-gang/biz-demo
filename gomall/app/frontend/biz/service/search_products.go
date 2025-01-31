package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/stigwang-gang/biz-demo/gomall/app/frontend/infra/rpc"
	rpcproduct "github.com/stigwang-gang/biz-demo/gomall/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/app"
	product "github.com/stigwang-gang/biz-demo/gomall/app/frontend/hertz_gen/frontend/product"
)

type SearchProductsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchProductsService(Context context.Context, RequestContext *app.RequestContext) *SearchProductsService {
	return &SearchProductsService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProductsService) Run(req *product.SearchProductsReq) (resp map[string]any, err error) {
	products, err := rpc.ProductClient.SearchProducts(h.Context, &rpcproduct.SearchProductsReq{
		Query: req.Q,
	})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"items": products.Results,
		"q":     req.Q,
	}, nil
}
