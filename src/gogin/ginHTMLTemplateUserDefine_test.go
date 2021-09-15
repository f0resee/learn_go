package gogin

import (
	"github.com/gin-gonic/gin"
	template2 "html/template"
	"net/http"
	"testing"
)

func TestUserDefineHTMLTemplate(t *testing.T)  {
	r := gin.Default()
	template := template2.Must(template2.ParseFiles("./templates/index.tmpl"))
	r.SetHTMLTemplate(template)
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.tmpl",gin.H{
			"title":"Main website",
		})
	})
	r.Run()
}
