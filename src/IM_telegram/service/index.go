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
	files, err := template.ParseFiles("src\\IM_telegram\\index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = files.Execute(c.Writer, "src\\IM_telegram\\index")
	if err != nil {
		fmt.Print(err)
		return
	}

	//c.JSON(200, gin.H{
	//	"message": "welcome",
	//})
}
