package gogin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestAsciiJson(t *testing.T)  {
	r :=gin.Default()
	r.GET("/someJson", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang":"GO语言\t哈哈",
			"tag":"/br",
		}
		//c.JSON(http.StatusOK,gin.H{
		//	"lang":"GO语言\t哈哈",
		//	"tag":"/br",
		//})
		c.AsciiJSON(http.StatusOK,data)

	})
	r.Run()
}