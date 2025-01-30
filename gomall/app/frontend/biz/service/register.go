package service

import (
	"context"
	"fmt"
	"github.com/hertz-contrib/sessions"
	"github.com/stigwang-gang/biz-demo/gomall/app/frontend/infra/rpc"
	"github.com/stigwang-gang/biz-demo/gomall/rpc_gen/kitex_gen/user"

	"github.com/cloudwego/hertz/pkg/app"
	auth "github.com/stigwang-gang/biz-demo/gomall/app/frontend/hertz_gen/frontend/auth"
	common "github.com/stigwang-gang/biz-demo/gomall/app/frontend/hertz_gen/frontend/common"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *auth.RegisterReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	userResp, err := rpc.UserClient.Register(h.Context, &user.RegisterReq{
		Email:           req.Email,
		Password:        req.Password,
		PasswordConfirm: req.PasswordConfirm,
	})
	fmt.Printf("userReq request: %+v\n", userResp)
	fmt.Printf("req request: %+v\n", req.PasswordConfirm)
	if err != nil {
		return nil, err
	}
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", userResp.UserId)
	err = session.Save()
	if err != nil {
		return nil, err
	}
	return
}
