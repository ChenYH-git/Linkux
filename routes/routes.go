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

	r.POST("/login", controllers.LoginHandler)
	r.GET("/search", controllers.SearchHandler)
	//r.LoadHTMLGlob("./template/*")
	//r.Static("/static", "./static")
	v1 := r.Group("/")
	v1.Use(middleware.JWTAuthMiddleware())
	{
		v1.GET("index", controllers.IndexHandler)
		v1.POST("post", controllers.CreatePostHandler)
		v1.POST("vote", controllers.PostVoteController)
		v1.PUT("view/add", controllers.AddViewHandler)
		v1.GET("label", controllers.LabelHandler)
		v1.GET("label/:id", controllers.LabelDetailHandler)
		v1.GET("post/:id", controllers.GetPostDetailHandler)
		v1.GET("rank", controllers.GetUserRankHandler)
		v1.GET("contribution", controllers.GetUserContributionHandler)
		v1.POST("collect", controllers.AddCollectionHandler)
		v1.GET("collect/get", controllers.GetCollectionHandler)
		v1.PUT("collect/delete", controllers.DeleteCollectionHandler)
		v1.POST("follow", controllers.AddFollowHandler)
		v1.GET("follow/get/follow", controllers.GetFollowUserHandler)
		v1.GET("follow/get/followed", controllers.GetFollowedUserHandler)
		v1.GET("follow/get/post", controllers.GetFollowPostHandler)
		v1.PUT("follow/cancel", controllers.CancelFollowHandler)
		v1.POST("trans", controllers.CreateTransTaskHandler)
		v1.GET("trans/get/task", controllers.GetTransTaskHandler)
		v1.GET("trans/get/exist", controllers.GetTransExistHandler)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "404 NOT FOUND",
		})
	})
	return r
}
