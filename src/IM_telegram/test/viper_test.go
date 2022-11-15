package test

import (
	"github.com/spf13/viper"
	"testing"
)

func TestViper(t *testing.T) {
	viper.SetConfigFile("F:\\go_project\\src\\IM_telegram\\config\\init.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return
	}
	t.Log(viper.GetString("mysql.dsn"))
}
