package service

import (
	"crypto/md5"
	"errors"
	"fmt"
	"gf-in-action/app/dao"
	"gf-in-action/app/model"
	"gf-in-action/library/utils"
)

var User = userService{}

type userService struct{}

var (
	ErrLoginFirst    = errors.New("请先登录")
	ErrLoginFailure  = errors.New("登录失败")
	ErrWrongPassword = errors.New("账户密码错误")
	ErrUserIsExist   = errors.New("用户名已存在")
)

// UserInfo 获取用户信息
func (u userService) UserInfo(uid int) (*model.User, error) {
	if uid <= 0 {
		return nil, nil
	}
	return dao.User.FindOne("id=?", uid)
}

// Login 用户登录
func (u userService) Login(username string, password string) (*model.User, error) {
	user, err := dao.User.FindOne("username=?", username)
	if err != nil {
		return nil, ErrLoginFailure
	}

	if user == nil || user.Id <= 0 {
		return nil, ErrWrongPassword
	}

	md5password := fmt.Sprintf("%x", md5.Sum([]byte(password+user.Passsalt)))
	if md5password != user.Password {
		return nil, ErrWrongPassword
	}

	return user, nil
}

// Register 用户注册
func (u userService) Register(username string, password string) error {
	if u.CheckUserNameExist(username) {
		return ErrUserIsExist
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
	user, err := dao.User.Fields("id").FindOne("username = ?", username)
	if err != nil || (user != nil && user.Id > 0) {
		return true
	}
	return false
}
