package dao

import (
	"room_rating_guidence/api/entity"
	"room_rating_guidence/common/util"
	"room_rating_guidence/infrs/db"
	"time"
)

// Register 注册
func Register(dto *entity.UserRegisterDto) (uint, error) {
	user := entity.User{
		Username:   dto.Username,
		Password:   util.EncryptionMd5(dto.Password),
		CreateTime: util.HTime{Time: time.Now()},
		UpdateTime: util.HTime{Time: time.Now()},
	}

	err := db.DB.Create(&user).Error
	return user.ID, err
}

func GetUserByUserName(username string) (user entity.User) {
	db.DB.Where("username = ?", username).First(&user)
	return user
}

func GetUserByUserId(id int) (user entity.User, err error) {
	err = db.DB.Where("id = ?", id).First(&user).Error
	return user, err
}
