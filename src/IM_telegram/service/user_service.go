package service

import (
	"github.com/gin-gonic/gin"
	"go_project/src/IM_telegram/models"
	"go_project/src/IM_telegram/utils"
	"strconv"
	"time"
)

// UserListHandler
// @Summary 返回用户列表
// @Schemes
// @Description 获得用户列表
// @Tags 用户模块
// @Accept json
// @Produce json
// @Success 200 {string} json{"code","message"}
// @Router /user/getuserlist [get]
func UserListHandler(c *gin.Context) {
	data := utils.GetUserList()
	c.JSON(200, gin.H{
		"message": data,
	})
}

// CreateUserHandler
// @Summary 创建用户
// @Schemes
// @Description 创建一个新用户
// @Tags 用户模块
// @Accept json
// @Produce json
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/createuser [get]
func CreateUserHandler(c *gin.Context) {
	user := models.UserBasic{
		LoginTime:    time.Now(),
		LoginOutTime: time.Now(),
		HeatBeatTime: time.Now(),
	}
	user.Name = c.Query("name")
	passwd := c.Query("password")
	repasswd := c.Query("repassword")
	if passwd != repasswd {
		c.JSON(200, gin.H{
			"message": "两次密码不一致",
		})
		return
	}
	user.Password = passwd
	utils.CreateUser(user)
	c.JSON(200, gin.H{
		"message": "新增用户成功!",
	})
}

// DeleteUserHandler
// @Summary 删除用户
// @Schemes
// @Description 删除一个用户
// @Tags 用户模块
// @Accept json
// @Produce json
// @param id query string false "ID"
// @Success 200 {string} json{"code","message"}
// @Router /user/deleteuser [get]
func DeleteUserHandler(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	utils.DeleteUser(user)
	c.JSON(200, gin.H{
		"message": "用户删除成功!",
	})
}

// UpdateUserHandler
// @Summary 修改用户信息
// @Description 修改一个用户的信息
// @Tags 用户模块
// @param id formData string false "ID"
// @param name formData string false "用户名"
// @param password formData string false "密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/updateuser [post]
func UpdateUserHandler(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.Password = c.PostForm("password")
	utils.UpdateUser(user)
	c.JSON(200, gin.H{
		"message": "用户修改成功!",
	})
}
