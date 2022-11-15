package service

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go_project/src/IM_telegram/models"
	"go_project/src/IM_telegram/utils"
	"math/rand"
	"net/http"
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
// @Description 创建一个新用户
// @Tags 用户模块
// @param name formData string false "用户名"
// @param phone formData string false "绑定手机号"
// @param email formData string false "注册邮箱"
// @param password formData string false "密码"
// @param repassword formData string false "确认密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/createuser [post]
func CreateUserHandler(c *gin.Context) {
	user := models.UserBasic{
		LoginTime:    time.Now(),
		LoginOutTime: time.Now(),
		HeatBeatTime: time.Now(),
	}
	user.Name = c.PostForm("name")
	//user.Phone = c.PostForm("phone")
	//user.Email = c.PostForm("email")
	passwd := c.PostForm("password")
	repasswd := c.PostForm("repasswd")
	if user.Name == "" || passwd == "" || repasswd == "" {
		c.JSON(200, gin.H{
			"code":    -1, //0成功 -1失败
			"message": "用户名或密码不能为空",
			"data":    nil,
		})
		return
	}
	if passwd != repasswd {
		c.JSON(200, gin.H{
			"code":    -1, //0成功 -1失败
			"message": "两次密码不一致",
			"data":    nil,
		})
		return
	}
	sqlUser := utils.FindUserByName(user.Name)
	if sqlUser.Name != "" {
		c.JSON(200, gin.H{
			"code":    -1, //0成功 -1失败
			"message": "用户名已存在",
			"data":    nil,
		})
		return
	}

	//sqlUser = utils.FindUserByPhone(user.Phone)
	//if sqlUser.Name != "" {
	//	c.JSON(200, gin.H{
	//		"code":    -1, //0成功 -1失败
	//		"message": "手机号码已注册",
	//		"data":    nil,
	//	})
	//	return
	//}
	//sqlUser = utils.FindUserByEmail(user.Email)
	//if sqlUser.Name != "" {
	//	c.JSON(200, gin.H{
	//		"code":    -1, //0成功 -1失败
	//		"message": "邮箱已注册",
	//		"data":    nil,
	//	})
	//	return
	//}

	//_, err := govalidator.ValidateStruct(user)
	//if err != nil {
	//	fmt.Println(err)
	//	c.JSON(200, gin.H{
	//		"code":    -1, //0成功 -1失败
	//		"message": "手机号码或邮箱格式不正确",
	//		"data":    nil,
	//	})
	//	return
	//}

	user.Salt = fmt.Sprintf("%d", rand.Int31())
	user.Password = utils.PassWdCrpyto(passwd, user.Salt)
	utils.CreateUser(user)
	c.JSON(200, gin.H{
		"code":    0, //0成功 -1失败
		"message": "新增用户成功!",
		"data":    user,
	})
}

// DeleteUserHandler
// @Summary 删除用户
// @Description 删除一个用户
// @Tags 用户模块
// @param id formData string false "ID"
// @Success 200 {string} json{"code","message"}
// @Router /user/deleteuser [post]
func DeleteUserHandler(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	utils.DeleteUser(user)
	c.JSON(200, gin.H{
		"code":    0, //0成功 -1失败
		"message": "用户删除成功!",
		"data":    nil,
	})
}

// UpdateUserHandler
// @Summary 修改用户信息
// @Description 修改一个用户的信息
// @Tags 用户模块
// @param id formData string false "ID"
// @param name formData string false "用户名"
// @param password formData string false "密码"
// @param phone formData string false "绑定手机号"
// @param email formData string false "绑定邮箱"
// @Success 200 {string} json{"code","message"}
// @Router /user/updateuser [post]
func UpdateUserHandler(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.Password = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")

	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"code":    -1, //0成功 -1失败
			"message": "用户修改失败!",
			"data":    nil,
		})
	} else {
		utils.UpdateUser(user)
		c.JSON(200, gin.H{
			"code":    0, //0成功 -1失败
			"message": "用户修改成功!",
			"data":    user,
		})
	}

}

// UserLoginHandler
// @Summary 用户登录
// @Description 用户登录
// @Tags 用户模块
// @param name formData string false "用户名"
// @param password formData string false "密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/login [post]
func UserLoginHandler(c *gin.Context) {

	name := c.PostForm("name")
	password := c.PostForm("password")

	sqlUserFindByName := utils.FindUserByName(name)
	if sqlUserFindByName.Name == "" {
		c.JSON(200, gin.H{
			"code":    -1, //0成功 -1失败
			"message": "用户名不正确!",
			"data":    nil,
		})
		return
	}

	flag := utils.ValidPassWd(password, sqlUserFindByName.Salt, sqlUserFindByName.Password)
	if !flag {
		c.JSON(200, gin.H{
			"code":    -1, //0成功 -1失败
			"message": "密码错误",
			"data":    nil,
		})
		return
	}

	sqlUserFindByNameAndPasswd := utils.FindUserByNameAndPasswd(name, sqlUserFindByName.Password)
	c.JSON(200, gin.H{
		"code":    0, //0成功 -1失败
		"message": "登录成功",
		"data":    sqlUserFindByNameAndPasswd,
	})
}

//防止跨域站点伪造请求
var upGrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendMsgHandler(c *gin.Context) {
	conn, err := upGrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(conn)
	MsgHandler(conn, c)
}

func MsgHandler(conn *websocket.Conn, c *gin.Context) {
	for {
		subscribe, err := utils.Subscribe(c, utils.PublishKey)
		if err != nil {
			return
		}
		tm := time.Now().Format("2006-01-02 15:04:05")
		msg := fmt.Sprintf("[ws][%s]:%s", tm, subscribe)
		fmt.Println(msg)
		err = conn.WriteMessage(1, []byte(msg))
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}

func SendUserMesg(c *gin.Context) {
	models.Chat(c.Writer, c.Request)
}
