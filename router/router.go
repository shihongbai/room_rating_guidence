package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"room_rating_guidence/api/controller"
	"room_rating_guidence/infrs/middleware"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	// 恢复
	router.Use(gin.Recovery())
	// 跨域请求中间件
	router.Use(middleware.Cors())
	// 图片访问路径静态文件夹可直接访问
	// 告诉gin框架模板文件引用的静态文件去哪里找
	router.Static("/static", "static")
	// 告诉gin框架去哪里找模板文件
	router.LoadHTMLGlob("templates/*")
	//router.StaticFS(config.Config.ImageSettings.UploadDir,
	//	http.Dir(config.Config.ImageSettings.UploadDir))
	// 日志log中间件
	router.Use(middleware.Logger())
	register(router)
	return router
}

// register 注册路由
func register(router *gin.Engine) {

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	router.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil)
	})

	// user相关路由
	user := router.Group("/api/user")
	{
		user.POST("/register", controller.Register)
		user.POST("/login", controller.Login)
	}
}
