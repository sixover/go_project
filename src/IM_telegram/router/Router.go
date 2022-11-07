package router

import (
	"github.com/gin-gonic/gin"
	"go_project/src/IM_telegram/service"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.GET("/index", service.GetIndex)
	router.GET("/user/getuserlist", service.UserList)
	return router
}
