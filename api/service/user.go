package service

import (
	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	"room_rating_guidence/api/dao"
	"room_rating_guidence/api/entity"
	"room_rating_guidence/common/result"
	"room_rating_guidence/common/util"
)

// IUserService 用户类接口
type IUserService interface {
	Register(c *gin.Context, dto *entity.UserRegisterDto)
	Login(c *gin.Context, dto *entity.UserLoginDto)
}

type UserServiceImpl struct{}

func NewUserServiceImpl() IUserService {
	return &UserServiceImpl{}
}

// Register 注册
func (u *UserServiceImpl) Register(c *gin.Context, dto *entity.UserRegisterDto) {
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.REQUIRED), result.ApiCode.GetMessage(result.ApiCode.REQUIRED))
		return
	}

	if dto.Password != dto.ConfirmPassword {
		result.Failed(c, int(result.ApiCode.FAILED), "两次密码不一致")
		return
	}

	// 检查是否用户已经存在
	user := dao.GetUserByUserName(dto.Username)
	if user.ID > 0 {
		result.Failed(c, int(result.ApiCode.FAILED), "用户已经存在")
		return
	}

	_, err = dao.Register(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.FAILED), "注册失败，请联系管理员")
		return
	}

	result.Success(c, "注册成功")
}

// Login 登录
func (u *UserServiceImpl) Login(c *gin.Context, dto *entity.UserLoginDto) {
	// 入参校验
	err := validator.New().Struct(dto)
	if err != nil {
		result.Failed(c, int(result.ApiCode.REQUIRED), result.ApiCode.GetMessage(result.ApiCode.REQUIRED))
	}

	user := dao.GetUserByUserName(dto.Username)
	if user.ID == 0 {
		result.Failed(c, int(result.ApiCode.FAILED), "用户名和秘密不正确")
		return
	}
	// 判断是否存在
	if user.Password != util.EncryptionMd5(dto.Password) {
		result.Failed(c, int(result.ApiCode.FAILED), "用户名和秘密不正确")
		return
	}
	result.Success(c, user)
}
