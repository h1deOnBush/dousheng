package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github/h1deOnBush/dousheng/internal/routers/api/v1"
	controller "github/h1deOnBush/dousheng/internal/controller"
)

func InitRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	user := v1.NewUser()
	video := v1.NewVideo()

	apiRouter := r.Group("/douyin")


	// basic apis
	apiRouter.GET("/feed/", video.Feed)
	apiRouter.GET("/user/", user.UserInfo)
	apiRouter.POST("/user/register/", user.Register)
	apiRouter.POST("/user/login/", user.Login)
	apiRouter.POST("/publish/action/", video.Publish)
	apiRouter.GET("/publish/list/", video.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", controller.FavoriteList)
	apiRouter.POST("/comment/action/", controller.CommentAction)
	apiRouter.GET("/comment/list/", controller.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", controller.FollowList)
	apiRouter.GET("/relation/follower/list/", controller.FollowerList)
}
