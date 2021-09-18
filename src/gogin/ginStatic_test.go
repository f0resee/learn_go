package gogin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestGinStatic(t *testing.T)  {
	r := gin.Default()
	r.StaticFile("/favicon.ico","./favicon.ico")
	r.LoadHTMLGlob("./templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.tmpl",gin.H{
			"title":"Main website",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}