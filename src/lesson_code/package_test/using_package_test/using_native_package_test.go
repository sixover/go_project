package using_package_test

import (
	"github.com/gin-gonic/gin"
	"go_project/src/lesson_code/package_test/src_package"
	"net/http"
	"testing"
)

func TestUsingNativeFibonaci(t *testing.T) {
	list, err := src_package.Fibonaci(10)
	if err != nil {
		t.Log("its an error")
		return
	}
	t.Log(list)
}

func TestNetPackage(t *testing.T) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
