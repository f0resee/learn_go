package gogin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestHTMLTemplate(t *testing.T)  {
	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.tmpl",gin.H{
			"title":"Main website",
		})
	})
	r.Run()
}
