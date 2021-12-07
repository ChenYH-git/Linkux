package routes

import (
	"Linkux/controllers"
	_ "Linkux/docs"
	"Linkux/logger"
	"Linkux/middleware"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	//初始化gin框架内置的校验器使用的翻译器
	controllers.InitTrans("zh")

	r := gin.New()

	r.Use(logger.GinLogger(), logger.GinRecovery(true), middleware.RateLimitMiddleware(60*time.Millisecond, 500)) // 令牌桶容量为500，每6秒钟填充100个

	r.LoadHTMLGlob("./template/*")
	r.Static("/static", "./static")

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/login", controllers.LoginHandler)
	r.GET("/search", controllers.SearchHandler)

	v1 := r.Group("/")
	v1.Use(middleware.JWTAuthMiddleware())
	{
		v1.GET("index", controllers.IndexHandler)
		v1.POST("post", controllers.CreatePostHandler)
		v1.POST("vote", controllers.PostVoteController)
		v1.POST("getvc", controllers.GetVoteCollectHandler)
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

	v2 := r.Group("/administer")
	v2.POST("/login", controllers.AdministerLoginHandler)
	v2.Use(middleware.AdministerCheck())
	{
		v2.GET("/examine/getposts", controllers.ExamineGetPostsHandler)
		v2.GET("/examine/gettask", controllers.ExamineGetTaskHandler)
		v2.PUT("/examine/put", controllers.ExaminePutChangesHandler)
		v2.GET("/posts/get", controllers.GetPostsExistsHandler)
		v2.PUT("/posts/star", controllers.StarPostsHandler)
		v2.PUT("/posts/star/cancel", controllers.CancelStarPostsHandler)
		v2.DELETE("/posts/delete", controllers.DeletePostsHandler)
		v2.DELETE("/trans/delete", controllers.DeleteTransHandler)
		v2.POST("/getp", controllers.GetPostStatusHandler)
		v2.GET("/user/get", controllers.GetUserExistsHandler)
		v2.PUT("/user/star", controllers.StarUserHandler)
		v2.PUT("/user/star/cancel", controllers.CancelStarUserHandler)
		v2.POST("/getu", controllers.GetUserStatusHandler)
	}

	v3 := r.Group("/page")
	{
		v3.GET("/login", func(c *gin.Context) {
			c.HTML(http.StatusOK, "login.html", nil)
		})
		v3.GET("/user", func(c *gin.Context) {
			c.HTML(http.StatusOK, "user.html", nil)
		})
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "404 NOT FOUND",
		})
	})
	return r
}
