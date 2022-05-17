package v1

import (
	"github.com/gin-gonic/gin"
	global "github/h1deOnBush/dousheng/gloabal"
	"github/h1deOnBush/dousheng/internal/service"
	"github/h1deOnBush/dousheng/pkg/app"
	"github/h1deOnBush/dousheng/pkg/convert"
	"github/h1deOnBush/dousheng/pkg/errcode"
)

type Relation struct {
}

func NewRelation() Relation {
	return Relation{}
}

type RelationActionRequest struct {
	UserId     int64 `json:"user_id" binding:"gt=0"`
	ToUserId   int64 `json:"to_user_id" binding:"gt=0"`
	ActionType int   `json:"action_type" binding:"oneof=1 2"`
}

type FollowListRequest struct {
	UserId int64 `json:"user_id" binding:"gt=0"`
}

type FollowListResponse struct {
	Response
	UserList []*service.User `json:"user_list"`
}

func (r Relation) RelationAction(c *gin.Context) {
	userIdI, _ := c.Get("user_id")
	userId := userIdI.(int64)
	req := RelationActionRequest{
		UserId:     userId,
		ToUserId:   convert.MustInt64(c.Query("to_user_id")),
		ActionType: convert.MustInt(c.Query("action_type")),
	}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &req)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())

	if req.ActionType == 1 {
		// 关注
		err := svc.Follow(req.ToUserId, req.UserId)
		if err != nil {
			global.Logger.Errorf(c, "service.Follow errs:%v", err)
			response.ToErrorResponse(errcode.FollowFail)
			return
		}
		response.ToResponse(Response{
			StatusCode: 0,
			StatusMsg:  "关注成功",
		})
	} else {
		// 取关
		err := svc.Unfollow(req.ToUserId, req.UserId)
		if err != nil {
			global.Logger.Errorf(c, "service.Unfollow errs:%v", err)
			response.ToErrorResponse(errcode.UnFollowFail)
			return
		}
		response.ToResponse(Response{
			StatusCode: 0,
			StatusMsg:  "取消关注成功",
		})
	}
}

func (r Relation) FollowList(c *gin.Context) {
	// 参数校验
	req := FollowListRequest{
		UserId: convert.MustInt64(c.Query("user_id")),
	}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &req)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())

	userList, err := svc.GetFollowList(req.UserId)
	if err != nil {
		global.Logger.Errorf(c, "service.GetFollowList err:%v", err)
		response.ToErrorResponse(errcode.GetFollowListFail)
		return
	}
	response.ToResponse(FollowListResponse{
		Response: Response{0, ""},
		UserList: userList,
	})
}

func (r Relation) FollowerList(c *gin.Context) {
	// 参数校验
	req := FollowListRequest{
		UserId: convert.MustInt64(c.Query("user_id")),
	}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &req)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())

	userList, err := svc.GetFollowerList(req.UserId)
	if err != nil {
		global.Logger.Errorf(c, "service.GetFollowList err:%v", err)
		response.ToErrorResponse(errcode.GetFollowerListFail)
		return
	}
	response.ToResponse(FollowListResponse{
		Response: Response{0, ""},
		UserList: userList,
	})
}
