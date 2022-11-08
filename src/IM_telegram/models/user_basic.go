package models

import (
	"gorm.io/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name         string
	Password     string
	Phone        string `valid:"matches((13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\\d{8})"`
	Email        string `valid:"email"`
	Identity     string
	ClientIP     string
	ClientPort   string
	LoginTime    time.Time
	LoginOutTime time.Time
	HeatBeatTime time.Time
	DeviceInfo   string
	IsLogOut     bool
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
