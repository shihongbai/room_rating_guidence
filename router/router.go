package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"room_rating_guidence/api/controller"
	"room_rating_guidence/common/config"
	"room_rating_guidence/infrs/middleware"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	// 恢复
	router.Use(gin.Recovery())
	// 跨域请求中间件
	router.Use(middleware.Cors())
	// 图片访问路径静态文件夹可直接访问
	router.StaticFS(config.Config.ImageSettings.UploadDir,
		http.Dir(config.Config.ImageSettings.UploadDir))
	// 日志log中间件
	router.Use(middleware.Logger())
	register(router)
	return router
}

// register 注册路由
func register(router *gin.Engine) {

	// user相关路由
	user := router.Group("/api/user")
	{
		user.POST("/register", controller.Register)
		user.POST("/login", controller.Login)
	}
}
