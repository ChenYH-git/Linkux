package routes

import (
	"Linkux/controllers"
	"Linkux/logger"
	"Linkux/middleware"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	//初始化gin框架内置的校验器使用的翻译器
	controllers.InitTrans("zh")

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true), middleware.RateLimitMiddleware(time.Second, 5)) // 令牌桶容量为5，每秒钟填充1个

	//r.LoadHTMLGlob("./template/*")
	//r.Static("/static", "./static")
	{
		r.POST("/login", controllers.LoginHandler)
		r.GET("/index", controllers.IndexHandler)
		r.POST("/post", controllers.CreatePostHandler)
		r.POST("/vote", controllers.PostVoteController)
		r.PUT("/view/add", controllers.AddViewHandler)
		r.GET("/label", controllers.LabelHandler)
		r.GET("/label/:id", controllers.LabelDetailHandler)
		r.GET("/post/:id", controllers.GetPostDetailHandler)
		r.GET("/rank", controllers.GetUserRankHandler)
		r.GET("/contribution", controllers.GetUserContributionHandler)
		r.POST("/collect", controllers.AddCollectionHandler)
		r.GET("/collect/get", controllers.GetCollectionHandler)
		r.PUT("/collect/delete", controllers.DeleteCollectionHandler)
		r.POST("/follow", controllers.AddFollowHandler)
		r.GET("/follow/get/follow", controllers.GetFollowUserHandler)
		r.GET("/follow/get/followed", controllers.GetFollowedUserHandler)
		r.GET("/follow/get/post", controllers.GetFollowPostHandler)
		r.PUT("/follow/cancel", controllers.CancelFollowHandler)
		r.POST("/trans", controllers.CreateTransTaskHandler)
		r.GET("/trans/get/task", controllers.GetTransTaskHandler)
		r.GET("/trans/get/exist", controllers.GetTransExistHandler)
		//r.POST("/search", controllers.SearchHandler)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "404 NOT FOUND",
		})
	})
	return r
}
