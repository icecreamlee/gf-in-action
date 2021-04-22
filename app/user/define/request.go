package define

type UserLoginReq struct {
	Username string `json:"username" v:"username@required|length:6,30#请输入用户名称|用户名称长度非法"`
	Password string `json:"password" v:"password@password"`
}

type UserRegisterReq struct {
	Username  string `json:"username" v:"username@required|length:6,30#请输入用户名称|用户名称长度非法"`
	Password  string `json:"password" v:"password@password|same:password2"`
	Password2 string `json:"password2" v:"password2@password|same:password#||两次密码不一致，请重新输入"`
}
