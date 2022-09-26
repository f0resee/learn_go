package hystrix

import (
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
	"time"
)

func TestHystrixInAServer(t *testing.T) {
	//http.ListenAndServe(":8090", &Handle{})
	hystrix.ConfigureCommand("mycommand", hystrix.CommandConfig{
		Timeout:                int(3 * time.Second),
		MaxConcurrentRequests:  10,
		SleepWindow:            5000,
		RequestVolumeThreshold: 20,
		ErrorPercentThreshold:  30,
	})
	//
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		msg := "success"
		_ = hystrix.Do("mycommand", func() error {
			_, err := http.Get("https://www.baidu.com")
			if err != nil {
				fmt.Printf("请求失败:%v", err)
				return err
			}
			return nil

		}, func(err error) error {
			fmt.Printf("handle  error:%v\n", err)
			msg = "error"
			return nil
		})

		if msg != "success" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": msg,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg": msg,
			})
		}
	})
	r.Run(":8090")
}
