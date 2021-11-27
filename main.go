package main

import (
	"Linkux/dao/mysql"
	"Linkux/dao/redis"
	"Linkux/logger"
	"Linkux/pkg/snowflakes"
	"Linkux/routes"
	"Linkux/settings"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

// @title 2021 秋软工团队合作项目
// @version 1.0
// @description 软工合作 alpha 冲刺版
// @termsOfService http://swagger.io/terms/

// @contact.name FZU UNDEFINED 小组
// @contact.url https://blog.csdn.net/codingyuhan
// @contact.email 2135378647@qq.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 119.23.175.25:8080
// @BasePath
func main() {
	//1.加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed, err:%v\n", err)
		return
	}
	//2.初始化日志
	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")
	//3.初始化mysql连接
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close()
	//4.初始化redis连接
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close()

	if err := snowflakes.Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
		fmt.Println("Init snowflakes failed")
		return
	}
	//5.注册路由
	r := routes.Setup(settings.Conf.Mode)
	//6.启动服务(优雅关机)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Conf.Port),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}
