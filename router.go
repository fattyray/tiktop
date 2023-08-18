package main

import (
	"github.com/gin-gonic/gin"
	"tiktop/controller"
	"tiktop/util"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	auth := util.Jwt2r()

	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.GET("/user/", auth, controller.UserInfo)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
	apiRouter.POST("/publish/action/", auth, controller.Publish)
	apiRouter.GET("/publish/list/", auth, controller.PublishList)

	apiRouter.Use(auth)
	{
		apiRouter.POST("/favorite/action/", controller.FavoriteAction)
		apiRouter.GET("/favorite/list/", controller.FavoriteList)
		apiRouter.POST("/comment/action/", controller.CommentAction)
		apiRouter.GET("/comment/list/", controller.CommentList)

		apiRouter.POST("/relation/action/", controller.RelationAction)
		apiRouter.GET("/relation/follow/list/", controller.FollowList)
		apiRouter.GET("/relation/follower/list/", controller.FollowerList)
		//apiRouter.GET("/relation/friend/list/", controller.FriendList)
		//apiRouter.GET("/message/chat/", controller.MessageChat)
		//apiRouter.POST("/message/action/", controller.MessageAction)
	}
}
