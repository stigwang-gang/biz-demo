package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/utils"
	//"github.com/stigwang-gang/biz-demo/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/stigwang-gang/biz-demo/gomall/app/frontend/infra/rpc"
	"github.com/stigwang-gang/biz-demo/gomall/rpc_gen/kitex_gen/product"

	"github.com/cloudwego/hertz/pkg/app"
	category "github.com/stigwang-gang/biz-demo/gomall/app/frontend/hertz_gen/frontend/category"
	//common "github.com/stigwang-gang/biz-demo/gomall/app/frontend/hertz_gen/frontend/common"
)

type CategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCategoryService(Context context.Context, RequestContext *app.RequestContext) *CategoryService {
	return &CategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *CategoryService) Run(req *category.CategoryReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	p, err := rpc.ProductClient.ListProducts(h.Context, &product.ListProductsReq{CategoryName: req.Category})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"title": "Category",
		"items": p.Products,
	}, nil
}
