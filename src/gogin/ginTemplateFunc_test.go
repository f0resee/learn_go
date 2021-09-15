package gogin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"testing"
	"time"
)

func formatAsDate(t time.Time)string  {
	year, month, date := t.Date()
	return fmt.Sprintf("%d/%02d/%02d",year,month,date)
}
func TestTemplateFunc(t *testing.T)  {
	r := gin.Default()
	r.Delims("{[{","}]}")
	r.SetFuncMap(template.FuncMap{
		"formatAsDate":formatAsDate,
	})
	r.LoadHTMLFiles("templates/raw.tmpl")
	r.GET("/date", func(c *gin.Context) {
		c.HTML(http.StatusOK,"raw.tmpl", map[string]interface{}{
			"now":time.Date(2017,0,0,0,0,0,0,time.UTC),
		})
	})
	r.Run()
}
