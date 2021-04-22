package api

import (
	"gf-in-action/app/user/define"
	"gf-in-action/app/user/service"
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
		response.Json(r, 1, err.Error())
		return
	}

	userRes := &define.UserInfoRes{}
	_ = copier.Copy(userRes, user)
	response.Json(r, 0, "success", g.Map{"user": userRes})
}

// Login 用户登录接口
func (u *userApi) Login(r *ghttp.Request) {
	req := &define.UserLoginReq{}
	_ = r.Parse(req)
	if e := gvalid.CheckStruct(req, nil); e != nil {
		response.Json(r, 1, e.FirstString())
		return
	}

	user, err := service.User.Login(req.Username, req.Password)
	if err != nil {
		response.Json(r, 1, err.Error())
		return
	}

	_ = r.Session.Set("uid", user.Id)
	response.Json(r, 0, "success")
}

// Register 用户注册接口
func (u *userApi) Register(r *ghttp.Request) {
	req := &define.UserRegisterReq{}
	_ = r.Parse(req)
	if e := gvalid.CheckStruct(req, nil); e != nil {
		response.Json(r, 1, e.FirstString(), req)
		return
	}

	err := service.User.Register(req.Username, req.Password)
	if err != nil {
		response.Json(r, 1, err.Error())
		return
	}
	response.Json(r, 0, "success")
}

// Logout 用户退出登录接口
func (u *userApi) Logout(r *ghttp.Request) {
	response.Json(r, 0, "logout")
}
