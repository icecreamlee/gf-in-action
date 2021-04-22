package service

import (
	"crypto/md5"
	"fmt"
	"gf-in-action/app/dao"
	"gf-in-action/app/model"
	"gf-in-action/app/user/define"
	"gf-in-action/library/utils"
	"github.com/gogf/gf/errors/gerror"
)

var User = userService{}

type userService struct{}

// UserInfo 获取用户信息
func (u userService) UserInfo(uid int) (*model.User, error) {
	if uid <= 0 {
		return nil, nil
	}
	return dao.User.FindOne(dao.User.Columns.Id+"=?", uid)
}

// Login 用户登录
func (u userService) Login(username string, password string) (*model.User, error) {
	user, err := dao.User.FindOne(dao.User.Columns.Username+"=?", username)
	if err != nil {
		return nil, gerror.NewCode(define.ErrCodeLoginFailure, define.ErrMsgLoginFailure)
	}

	if user == nil || user.Id <= 0 {
		return nil, gerror.NewCode(define.ErrCodeWrongPassword, define.ErrMsgWrongPassword)
	}

	md5password := fmt.Sprintf("%x", md5.Sum([]byte(password+user.Passsalt)))
	if md5password != user.Password {
		return nil, gerror.NewCode(define.ErrCodeWrongPassword, define.ErrMsgWrongPassword)
	}

	return user, nil
}

// Register 用户注册
func (u userService) Register(username string, password string) error {
	if u.CheckUserNameExist(username) {
		return gerror.NewCode(define.ErrCodeUserIsExist, define.ErrMsgUserIsExist)
	}

	passsalt := utils.GenRandomStr(6)
	md5password := fmt.Sprintf("%x", md5.Sum([]byte(password+passsalt)))

	_, err := dao.User.Insert(model.User{
		Username: username,
		Password: md5password,
		Passsalt: passsalt,
	})
	return err
}

// CheckUserNameExist 检查用户名是否已存在
func (u userService) CheckUserNameExist(username string) bool {
	user, err := dao.User.Fields(dao.User.Columns.Id).FindOne(dao.User.Columns.Username+"=?", username)
	if err != nil || (user != nil && user.Id > 0) {
		return true
	}
	return false
}
