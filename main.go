package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"room_rating_guidence/common/config"
	"room_rating_guidence/infrs/db"
	"room_rating_guidence/infrs/log"
	"room_rating_guidence/router"
	"time"
)

// @title 个人看房指南
// @ version 1.0
func main() {
	// 加载日志
	log := log.Log()
	// 设置启动模式
	gin.SetMode(config.Config.Server.Model)
	// 路由初始化
	router := router.InitRouter()
	srv := &http.Server{
		Addr:    config.Config.Server.Address,
		Handler: router,
	}
	// 启动服务
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Infof("listen: %s \n", err)
		}
		log.Infof("listen: %s \n", config.Config.Server.Address)
	}()
	quit := make(chan os.Signal, 1)
	// 监听消息
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Info("Server Shutdown:", err)
	}
	log.Info("Server exiting")
}

// 初始化连接
func init() {
	// mysql
	db.SetupDBLink()
	// redis
	//redis.SetupRedisDb()
}
