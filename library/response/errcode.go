package response

import "github.com/gogf/gf/errors/gerror"

const (
	ErrCodeSuccess = 0
	ErrMsgSuccess  = "操作成功"
	ErrCodeDefault = 1000
	ErrMsgDefault  = "操作失败"
)

// ErrCode 获取错误码
func ErrCode(err error) int {
	if err == nil {
		return ErrCodeSuccess
	}

	code := gerror.Code(err)
	if code == -1 {
		return ErrCodeDefault
	}

	return code
}

// ErrMsg 获取错误描述
func ErrMsg(err error) string {
	if err == nil {
		return ErrMsgSuccess
	}
	return err.Error()
}
