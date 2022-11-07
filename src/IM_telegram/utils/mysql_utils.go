package utils

import "go_project/src/IM_telegram/models"

func GetUserList() []*models.UserBasic {
	data := make([]*models.UserBasic, 10)
	DB.Find(&data)
	return data
}
