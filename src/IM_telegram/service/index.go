package service

import "github.com/gin-gonic/gin"

// GetIndex
// @Summary index example
// @Schemes
// @Description visit index
// @Tags 首页
// @Accept json
// @Produce json
// @Success 200 {string} Welcome
// @Router /index [get]
func GetIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "welcome",
	})
}
