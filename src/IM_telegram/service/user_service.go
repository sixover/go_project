package service

import (
	"github.com/gin-gonic/gin"
	"go_project/src/IM_telegram/utils"
)

func UserList(c *gin.Context) {
	data := utils.GetUserList()
	c.JSON(200, gin.H{
		"message": data,
	})
}
