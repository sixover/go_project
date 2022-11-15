package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go_project/src/IM_telegram/docs"
	"go_project/src/IM_telegram/service"
)

func Router() *gin.Engine {

	router := gin.Default()

	docs.SwaggerInfo.BasePath = ""
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Static("/asset", "src\\IM_telegram\\asset\\")
	router.LoadHTMLGlob("src\\IM_telegram\\views\\**\\*")

	router.GET("/", service.GetIndex)
	router.GET("/index", service.GetIndex)
	router.GET("/toRegister", service.ToRegister)

	router.GET("/user/getuserlist", service.UserListHandler)
	router.POST("/user/createuser", service.CreateUserHandler)
	router.POST("/user/deleteuser", service.DeleteUserHandler)
	router.POST("/user/updateuser", service.UpdateUserHandler)
	router.POST("/user/login", service.UserLoginHandler)

	router.GET("/user/sendMsg", service.SendMsgHandler)
	router.GET("/user/SendUserMesg", service.SendUserMesg)

	return router
}
