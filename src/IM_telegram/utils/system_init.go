package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitConfig() {
	viper.SetConfigName("init")
	viper.AddConfigPath("src\\IM_telegram\\config")
	err := viper.ReadInConfig()
	if err != nil {
		return
	}
	fmt.Println(viper.Get("mysql"))
}

func InitMysql() {
	var err error
	DB, err = gorm.Open(mysql.Open(viper.GetString("mysql.dsn")), &gorm.Config{})
	if err != nil {
		return
	}
}
