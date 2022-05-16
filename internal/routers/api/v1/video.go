package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	global "github/h1deOnBush/dousheng/gloabal"
	"github/h1deOnBush/dousheng/internal/service"
	"github/h1deOnBush/dousheng/pkg/app"
	"github/h1deOnBush/dousheng/pkg/convert"
	"github/h1deOnBush/dousheng/pkg/errcode"
	"github/h1deOnBush/dousheng/pkg/upload"
	"time"
)

type Video struct {
}

func NewVideo() Video {
	return Video{}
}

type PublishListResponse struct {
	Response
	VideoList []*service.Video `json:"video_list"`
}

type FeedResponse struct {
	Response
	NextTime  int64            `json:"next_time"`
	VideoList []*service.Video `json:"video_list"`
}

func (v Video) Publish(c *gin.Context) {
	// 参数校验
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("data")
	if err != nil {
		global.Logger.Errorf(c, "invalid param err:%v", err)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	svc := service.New(c.Request.Context())
	// 上传视频存储到本地
	fileInfo, err := svc.UploadFile(upload.TypeVideo, file, fileHeader)
	if err != nil {
		global.Logger.Errorf(c, "service.UploadFile err: %v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail)
		return
	}

	// 在数据库中创建一条记录
	userIdI, _ := c.Get("user_id")
	userId := userIdI.(int64)
	err = svc.Publish(userId, fileInfo.AccessUrl, "")

	if err != nil {
		global.Logger.Errorf(c, "dao.CreateVideo fail, err:%v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail)
		return
	}

	response.ToResponse(Response{
		StatusCode: 0,
		StatusMsg:  fmt.Sprintf("%v upload success", fileInfo.Name),
	})
}

func (v Video) PublishList(c *gin.Context) {
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())
	userIdI, _ := c.Get("user_id")
	userId := userIdI.(int64)

	videos, err := svc.PublishList(userId)
	if err != nil {
		global.Logger.Errorf(c, "service.PublishList, err:%v", err)
		response.ToErrorResponse(errcode.GetPublishListFail)
		return
	}

	response.ToResponse(PublishListResponse{
		Response:  Response{0, ""},
		VideoList: videos,
	})
}

func (v Video) Feed(c *gin.Context) {
	// 提取参数
	response := app.NewResponse(c)
	var t time.Time
	latestTime := c.Query("latest_time")
	if latestTime == "" {
		t = time.Now()
	} else {
		t = time.Unix(convert.MustInt64(latestTime[:len(latestTime)-3]), 0)
	}
	token := c.Query("token")
	var userId int64
	if token != "" {
		claims, err := app.ParseToken(token)
		if err == nil {
			userId = claims.Id
		}
	}
	var nextTime int64
	svc := service.New(c.Request.Context())
	videos, nextTime, err := svc.Feed(t, userId)
	if err != nil {
		global.Logger.Errorf(c, "Service.Feed error, err:%v", err)
		response.ToErrorResponse(errcode.GetFeedFail)
		return
	}

	response.ToResponse(FeedResponse{
		Response:  Response{0, ""},
		NextTime:  nextTime,
		VideoList: videos,
	})
}
