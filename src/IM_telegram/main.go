package main

import (
	"go_project/src/IM_telegram/router"
	"go_project/src/IM_telegram/utils"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	utils.InitConfig()
	utils.InitMysql()
	utils.InitRedis()

	r := router.Router()
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
