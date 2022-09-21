package gogin

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestGinFastExample(t *testing.T) {
	r := gin.Default()
	r.LoadHTMLFiles("./submitdata.html")
	r.GET("/submit", func(c *gin.Context) {
		fmt.Println("test")
		name := c.Query("name")
		if name != "" {
			fmt.Println(name)
		}
		c.HTML(http.StatusOK, "submitdata.html", gin.H{
			"title": name,
		})
	})
}

func withTimeout() func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second)
		c.Request = c.Request.WithContext(ctx)
		log.Println("processing")
		defer func() {
			if ctx.Err() == context.DeadlineExceeded {
				c.Abort()
			}
			cancel()
		}()

		c.Next()
	}
}

// 中间件实现超时返回
func TestGinWithTimeout(t *testing.T) {
	r := gin.Default()
	r.Use(withTimeout())
	r.GET("/ping", func(c *gin.Context) {
		ctx := c.Request.Context()
		res := make(chan int)
		go func() {
			time.Sleep(5 * time.Second)
			c.Writer.WriteHeader(http.StatusOK)
			res <- 1
		}()
		select {
		case <-ctx.Done():
			c.Writer.WriteHeader(http.StatusGatewayTimeout)
			return
		case <-res:
			return
		}
	})
	r.Run(":9000")
}

//直接实现
func TestGinWithTimeoutDirectly(t *testing.T) {
	r := gin.Default()
	r.Use(withTimeout())
	r.GET("/ping", func(c *gin.Context) {
		ctx := c.Request.Context()
		toCtx, _ := context.WithTimeout(ctx, 1*time.Second)
		c.Request = c.Request.WithContext(toCtx)
		res := make(chan int)
		go func(c1 context.Context) {
			select {
			case <-toCtx.Done():
				c.Writer.WriteHeader(http.StatusGatewayTimeout)
				return
			case <-res:
				time.Sleep(5 * time.Second)
				c.Writer.WriteHeader(http.StatusOK)
				res <- 1
				return
			}

		}(toCtx)

	})
	r.Run(":9000")
}
