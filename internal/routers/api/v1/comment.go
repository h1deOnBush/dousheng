package v1

import (
	"github.com/gin-gonic/gin"
	global "github/h1deOnBush/dousheng/gloabal"
	"github/h1deOnBush/dousheng/internal/service"
	"github/h1deOnBush/dousheng/pkg/app"
	"github/h1deOnBush/dousheng/pkg/convert"
	"github/h1deOnBush/dousheng/pkg/errcode"
)

type Comment struct {
}

func NewComment() Comment {
	return Comment{}
}

type CommentActionRequest struct {
	UserId  int64 `json:"user_id" binding:"gt=0"`
	VideoId int64 `json:"video_id" binding:"gt=0"`
	// 1发布评论 2删除
	ActionType  int    `json:"action_type" binding:"oneof=1 2"`
	CommentText string `json:"comment_text"`
	CommentId   int64  `json:"comment_id"`
}

type CommentListRequest struct {
	UserId  int64 `json:"user_id" binding:"gt=0"`
	VideoId int64 `json:"video_id" binding:"gt=0"`
}

type CommentListResponse struct {
	Response
	CommentList []*service.Comment `json:"comment_list"`
}

func (cmt Comment) CommentAction(c *gin.Context) {
	// 参数校验
	req := CommentActionRequest{
		UserId:      convert.MustInt64(c.Query("user_id")),
		VideoId:     convert.MustInt64(c.Query("video_id")),
		ActionType:  convert.MustInt(c.Query("action_type")),
		CommentText: c.Query("comment_text"),
		CommentId:   convert.MustInt64(c.Query("comment_id")),
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
		// 发布评论
		err := svc.CommentOn(req.UserId, req.VideoId, req.CommentText)
		if err != nil {
			global.Logger.Errorf(c, "service.CommentOn errs:%v", err)
			response.ToErrorResponse(errcode.CommentFail)
			return
		}
		response.ToResponse(Response{
			StatusCode: 0,
			StatusMsg:  "评论成功",
		})
	} else {
		// 删除评论
		err := svc.DeleteComment(req.CommentId, req.VideoId)
		if err != nil {
			global.Logger.Errorf(c, "service.DeleteComment errs:%v", err)
			response.ToErrorResponse(errcode.DeleteCommentFail)
			return
		}
		response.ToResponse(Response{
			StatusCode: 0,
			StatusMsg:  "删除评论成功",
		})
	}
}

func (cmt Comment) CommentList(c *gin.Context) {
	// 参数校验
	req := CommentListRequest{
		UserId:  convert.MustInt64(c.Query("user_id")),
		VideoId: convert.MustInt64(c.Query("video_id")),
	}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &req)
	if !valid {
		global.Logger.Errorf(c, "app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())

	commentList, err := svc.GetCommentList(req.UserId, req.VideoId)
	if err != nil {
		global.Logger.Errorf(c, "service.GetCommentList errs:%v", err)
		response.ToErrorResponse(errcode.GetCommentListFail)
		return
	}
	response.ToResponse(CommentListResponse{
		Response:    Response{0, ""},
		CommentList: commentList,
	})
}
