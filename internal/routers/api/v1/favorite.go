package v1

import (
	"github.com/gin-gonic/gin"
	global "github/h1deOnBush/dousheng/gloabal"
	"github/h1deOnBush/dousheng/internal/service"
	"github/h1deOnBush/dousheng/pkg/app"
	"github/h1deOnBush/dousheng/pkg/convert"
	"github/h1deOnBush/dousheng/pkg/errcode"
)

type Favorite struct {
}

func NewFavorite() Favorite {
	return Favorite{}
}

type FavoriteActionRequest struct {
	UserId  int64 `json:"user_id" binding:"gt=0"`
	VideoId int64 `json:"video_id" binding:"gt=0"`
	// 1代表点赞 2代表取消点赞
	ActionType int `json:"action_type" binding:"oneof=1 2"`
}

type FavoriteListRequest struct {
	UserId int64 `json:"user_id" binding:"gt=0"`
}

type FavoriteListResponse struct {
	Response
	VideoList []*service.Video `json:"video_list"`
}

func (f Favorite) FavoriteAction(c *gin.Context) {
	// 参数校验
	req := FavoriteActionRequest{
		UserId:     convert.MustInt64(c.Query("user_id")),
		VideoId:    convert.MustInt64(c.Query("video_id")),
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
		// 点赞
		err := svc.Like(req.UserId, req.VideoId)
		if err != nil {
			global.Logger.Errorf(c, "service.Like errs:%v", err)
			response.ToErrorResponse(errcode.LikeFail)
			return
		}
		response.ToResponse(Response{
			StatusCode: 0,
			StatusMsg:  "点赞成功",
		})
	} else {
		// 取消点赞
		err := svc.Unlike(req.UserId, req.VideoId)
		if err != nil {
			global.Logger.Errorf(c, "service.Unlike errs:%v", err)
			response.ToErrorResponse(errcode.UnlikeFail)
			return
		}
		response.ToResponse(Response{
			StatusCode: 0,
			StatusMsg:  "取消点赞成功",
		})
	}
}

func (f Favorite) FavoriteList(c *gin.Context) {
	// 参数校验
	req := FavoriteListRequest{
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
	videoList, err := svc.GetFavoriteList(req.UserId)
	if err != nil {
		global.Logger.Errorf(c, "service.GetFavoriteList errs:%v", err)
		response.ToErrorResponse(errcode.GetFavoriteListFail)
		return
	}

	response.ToResponse(FavoriteListResponse{
		Response:  Response{0, ""},
		VideoList: videoList,
	})
}
