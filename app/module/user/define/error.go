package define

import "gf-in-action/library/response"

// 错误码定义
//goland:noinspection GoUnusedGlobalVariable
var (
	// 基础错误码定义

	ErrCodeSuccess        = response.ErrCodeSuccess
	ErrMsgSuccess         = response.ErrMsgSuccess
	ErrCodeDefault        = response.ErrCodeDefault
	ErrMsgDefault         = response.ErrMsgDefault
	ErrCodeUnknown        = 1001
	ErrMsgUnknown         = "未知错误"
	ErrCodeFormValidation = 1002
	ErrMsgFormValidation  = "请求参数错误"

	// 业务错误码定义

	ErrCodeLoginFirst    = 10000
	ErrMsgLoginFirst     = "请先登录"
	ErrCodeLoginFailure  = 10001
	ErrMsgLoginFailure   = "登录失败"
	ErrCodeWrongPassword = 10002
	ErrMsgWrongPassword  = "账户密码错误"
	ErrCodeUserIsExist   = 10003
	ErrMsgUserIsExist    = "用户名已存在"
)
