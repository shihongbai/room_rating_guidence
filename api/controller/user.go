package controller

import (
	"github.com/gin-gonic/gin"
	"room_rating_guidence/api/entity"
	"room_rating_guidence/api/service"
)

func Register(c *gin.Context) {
	var dao entity.UserRegisterDto

	c.BindJSON(&dao)
	service.NewUserServiceImpl().Register(c, &dao)
}

func Login(c *gin.Context) {
	var dao entity.UserLoginDto

	c.BindJSON(&dao)
	service.NewUserServiceImpl().Login(c, &dao)
}
