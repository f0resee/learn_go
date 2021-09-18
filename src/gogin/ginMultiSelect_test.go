package gogin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

type myForm struct {
	Colors []string `form:"colors[]"`
}

func formHandler(c *gin.Context) {
	var fakeForm myForm
	c.ShouldBind(&fakeForm)
	c.JSON(200, gin.H{"color": fakeForm.Colors})
}

func TestGinMultiSelect(t *testing.T)  {
	r := gin.Default()
	r.LoadHTMLFiles("./form.html")
	r.GET("/form.html", func(c *gin.Context) {
		c.HTML(http.StatusOK,"form.html",gin.H{

		})
	})
	r.POST("/",formHandler)
	r.Run()
}