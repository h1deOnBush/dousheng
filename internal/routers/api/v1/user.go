package v1

import (
	"github.com/gin-gonic/gin"
	global "github/h1deOnBush/dousheng/gloabal"
	"github/h1deOnBush/dousheng/internal/service"
	"github/h1deOnBush/dousheng/pkg/app"
	"github/h1deOnBush/dousheng/pkg/convert"
	"github/h1deOnBush/dousheng/pkg/errcode"
)

type User struct{}

func NewUser() User {
	return User{}
}

type Response struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

// 用户注册登录请求
type RegisterRequest struct {
	Username string `binding:"max=32"`
	Password string `binding:"max=32"`
}

// 用户注册登录请求响应
type RegisterResponse struct {
	Response
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

// 用户认证登录请求
type LoginRequest struct {
	Username string `binding:"max=32"`
	Password string `binding:"max=32"`
}

// 用户认证响应
type LoginResponse struct {
	Response
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

// 获取用户信息请求
type GetUserInfoRequest struct {
	// 需要请求的用户信息id, 可能是自己，可能不是自己
	ToUserId int64
	// 自己的用户id
	UserId int64
}

type GetUserInfoResponse struct {
	Response
	User *service.User `json:"user"`
}

func (u User) Register(c *gin.Context) {
	// 参数校验
	req := RegisterRequest{
		Username: c.Query("username"),
		Password: c.Query("password"),
	}
	valid, errs := app.BindAndValid(c, &req)
	response := app.NewResponse(c)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 数据库操作
	svc := service.New(c.Request.Context())

	userId, err := svc.Register(req.Username, req.Password)
	if err != nil {
		global.Logger.Errorf(c, " service.Register:err:%v", err)
		response.ToErrorResponse(errcode.UserRegisterFail)
		return
	}

	// 生成token
	tokenString, err := app.GenToken(req.Username, userId)

	if err != nil {
		global.Logger.Errorf(c, "generate token fail:%v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	// 成功响应
	response.ToResponse(RegisterResponse{
		Response: Response{0, errcode.RegisterSuccess},
		UserId:   userId,
		Token:    tokenString,
	})
}

func (u User) Login(c *gin.Context) {
	// 参数校验
	req := LoginRequest{
		Username: c.Query("username"),
		Password: c.Query("password"),
	}
	valid, errs := app.BindAndValid(c, &req)
	response := app.NewResponse(c)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	// 数据库操作
	svc := service.New(c.Request.Context())

	userId, err := svc.CheckUser(req.Username, req.Password)
	if err != nil {
		global.Logger.Errorf(c, " service.Login:err:%v", err)
		response.ToErrorResponse(errcode.UserLoginFail)
		return
	}

	// 生成token
	tokenString, err := app.GenToken(req.Username, userId)

	if err != nil {
		global.Logger.Errorf(c, "generate token fail:%v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	// 成功响应
	response.ToResponse(LoginResponse{
		Response: Response{0, errcode.LoginSuccess},
		UserId:   userId,
		Token:    tokenString,
	})
}

func (u User) GetUserInfo(c *gin.Context) {
	uid, _ := c.Get("user_id")
	userId := uid.(int64)
	// 参数校验
	req := GetUserInfoRequest{
		ToUserId: convert.MustInt64(c.Query("user_id")),
		UserId:   userId,
	}
	valid, errs := app.BindAndValid(c, &req)
	response := app.NewResponse(c)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	// 数据库操作
	svc := service.New(c.Request.Context())
	user, err := svc.GetUserInfo(req.ToUserId, req.UserId)
	if err != nil {
		global.Logger.Errorf(c, "service.GetUserInfo fail:%v", err)
		response.ToErrorResponse(errcode.GetUserInfoFail)
		return
	}
	// 成功响应
	response.ToResponse(GetUserInfoResponse{
		Response: Response{0, ""},
		User:     user,
	})
}
