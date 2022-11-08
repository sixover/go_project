package utils

import (
	"go_project/src/IM_telegram/models"
	"gorm.io/gorm"
)

func GetUserList() []*models.UserBasic {
	data := make([]*models.UserBasic, 10)
	DB.Find(&data)
	return data
}

func FindUserByName(name string) models.UserBasic {
	user := models.UserBasic{}
	DB.Where("name = ?", name).First(&user)
	return user
}

func FindUserByPhone(phone string) models.UserBasic {
	user := models.UserBasic{}
	DB.Where("phone = ?", phone).First(&user)
	return user
}

func FindUserByEmail(email string) models.UserBasic {
	user := models.UserBasic{}
	DB.Where("email = ?", email).First(&user)
	return user
}

func CreateUser(user models.UserBasic) *gorm.DB {
	return DB.Create(&user)
}

func DeleteUser(user models.UserBasic) *gorm.DB {
	return DB.Delete(&user)
}

//func UpdateUser(user models.UserBasic) *gorm.DB {
//	return DB.Model(&user).Updates(models.UserBasic{
//		Model:        user.Model,
//		Name:         user.Name,
//		Password:     user.Password,
//		Phone:        user.Phone,
//		Email:        user.Email,
//		Identity:     user.Identity,
//		ClientIP:     user.ClientIP,
//		ClientPort:   user.ClientPort,
//		DeviceInfo:   user.DeviceInfo,
//		IsLogOut:     user.IsLogOut,
//	})
//}
func UpdateUser(user models.UserBasic) *gorm.DB {
	return DB.Model(&user).Updates(models.UserBasic{
		Name:     user.Name,
		Password: user.Password,
		Phone:    user.Phone,
		Email:    user.Email,
	})
}
