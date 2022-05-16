package routers

import (
	"github.com/gin-gonic/gin"
	"github/h1deOnBush/dousheng/internal/middlewares"
	v1 "github/h1deOnBush/dousheng/internal/routers/api/v1"
)

func InitRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	user := v1.NewUser()
	//video := v1.NewVideo()

	apiRouter := r.Group("/douyin")
	// basic apis
	//apiRouter.GET("/feed/", video.Feed)
	apiRouter.GET("/user/", middlewares.JWT(), user.GetUserInfo)
	apiRouter.POST("/user/register/", user.Register)
	apiRouter.POST("/user/login/", user.Login)
	//apiRouter.POST("/publish/action/", middlewares.JWT(), video.Publish)
	//apiRouter.GET("/publish/list/", middlewares.JWT(), video.PublishList)

	//// extra apis - I
	//apiRouter.POST("/favorite/action/", service.FavoriteAction)
	//apiRouter.GET("/favorite/list/", service.FavoriteList)
	//apiRouter.POST("/comment/action/", service.CommentAction)
	//apiRouter.GET("/comment/list/", service.CommentList)
	//
	//// extra apis - II
	//apiRouter.POST("/relation/action/", service.RelationAction)
	//apiRouter.GET("/relation/follow/list/", service.FollowList)
	//apiRouter.GET("/relation/follower/list/", service.FollowerList)
}
