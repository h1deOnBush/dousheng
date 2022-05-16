package errcode

// 用户错误码
var (
	UsernameAlreadyExists = NewError(20000000, "用户名已存在")
	UserNotExists         = NewError(20000001, "用户不存在")
	UserPasswordWrong     = NewError(20000002, "用户密码不存在")
	UserRegisterFail      = NewError(20000003, "用户注册失败")
	UserLoginFail         = NewError(20000004, "用户登录失败")
	GetUserInfoFail       = NewError(20000005, "获取用户信息失败")
)
