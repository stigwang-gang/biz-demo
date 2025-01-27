package service

import (
	"context"
	"github.com/hertz-contrib/sessions"

	"github.com/cloudwego/hertz/pkg/app"
	auth "github.com/stigwang-gang/biz-demo/gomall/app/frontend/hertz_gen/frontend/auth"
	//common "github.com/stigwang-gang/biz-demo/gomall/app/frontend/hertz_gen/frontend/common"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginReq) (resp string, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	//resp, err := rpc.UserClient.Login(h.Context, &user.LoginReq{
	//	Email:    req.Email,
	//	Password: req.Password,
	//})
	//if err != nil {
	//	return "", err
	//}

	session := sessions.Default(h.RequestContext)
	session.Set("user_id", 5)
	err = session.Save()
	if err != nil {
		return "", err
	}
	redirect := "/"
	if req.Next != "" {
		redirect = req.Next
	}

	return redirect, nil
}
