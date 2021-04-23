package api

import (
	"gf-in-action/app/module/user/define"
	"gf-in-action/app/module/user/service"
	"gf-in-action/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
	"github.com/jinzhu/copier"
)

var User = userApi{}

type userApi struct {
	Module string
}

// Index Hello World!
func (u *userApi) Index(r *ghttp.Request) {
	r.Response.Writeln("Hello World!")
}

// Info 查看用户信息接口
func (u userApi) Info(r *ghttp.Request) {
	uid := r.Session.GetInt("uid")
	user, err := service.User.UserInfo(uid)
	if err != nil {
		response.Json(r, define.ErrCodeDefault, err.Error())
		return
	}

	userRes := &define.UserInfoRes{}
	_ = copier.Copy(userRes, user)
	response.JsonWithError(r, nil, g.Map{"user": userRes})
}

// Login 用户登录接口
func (u *userApi) Login(r *ghttp.Request) {
	req := &define.UserLoginReq{}
	_ = r.Parse(req)
	if e := gvalid.CheckStruct(req, nil); e != nil {
		response.Json(r, define.ErrCodeFormValidation, e.FirstString())
		return
	}

	user, err := service.User.Login(req.Username, req.Password)
	if user != nil {
		_ = r.Session.Set("uid", user.Id)
	}

	response.JsonWithError(r, err)
}

// Register 用户注册接口
func (u *userApi) Register(r *ghttp.Request) {
	req := &define.UserRegisterReq{}
	_ = r.Parse(req)
	if e := gvalid.CheckStruct(req, nil); e != nil {
		response.Json(r, define.ErrCodeFormValidation, e.FirstString(), req)
		return
	}
	response.JsonWithError(r, service.User.Register(req.Username, req.Password))
}

// Logout 用户退出登录接口
func (u *userApi) Logout(r *ghttp.Request) {
	_ = r.Session.Clear()
	response.JsonWithError(r, nil)
}
