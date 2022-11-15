package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"text/template"
)

// GetIndex
// @Summary index example
// @Description visit index
// @Tags 首页
// @Success 200 {string} Welcome
// @Router /index [get]
func GetIndex(c *gin.Context) {
	files, err := template.ParseFiles("src\\IM_telegram\\index.html", "src\\IM_telegram\\views\\chat\\head.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = files.Execute(c.Writer, "index")
	if err != nil {
		fmt.Print(err)
		return
	}

	//c.JSON(200, gin.H{
	//	"message": "welcome",
	//})
}

func ToRegister(c *gin.Context) {
	files, err := template.ParseFiles("src\\IM_telegram\\views\\user\\register.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = files.Execute(c.Writer, "index")
	if err != nil {
		fmt.Print(err)
		return
	}
}
