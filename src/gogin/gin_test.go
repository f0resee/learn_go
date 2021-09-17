package gogin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
)

func TestGinFastExample(t *testing.T)  {
	r := gin.Default()
	r.LoadHTMLFiles("./submitdata.html")
	r.GET("/submit", func(c *gin.Context) {
		fmt.Println("test")
		name := c.Query("name")
		if name!=""{
			fmt.Println(name)
		}
		c.HTML(http.StatusOK,"submitdata.html",gin.H{
			"title":name,
		})
	})
	r.Run()
}