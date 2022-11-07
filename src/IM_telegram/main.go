package main

import (
	"go_project/src/IM_telegram/router"
	"go_project/src/IM_telegram/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMysql()

	r := router.Router()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
