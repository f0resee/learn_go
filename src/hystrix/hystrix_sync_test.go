package hystrix

import (
	"errors"
	"github.com/afex/hystrix-go/hystrix"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestHystrix(t *testing.T) {
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	go http.ListenAndServe(":8074", hystrixStreamHandler)

	hystrix.ConfigureCommand("aaa", hystrix.CommandConfig{
		Timeout:                1000,
		MaxConcurrentRequests:  1,
		SleepWindow:            5000,
		RequestVolumeThreshold: 1,
		ErrorPercentThreshold:  1,
	})

	for i := 0; i < 1000; i++ {
		err := hystrix.Do("aaa", func() error {
			if i == 0 {
				return errors.New("service error")
			}
			log.Println("do service")
			return nil
		}, nil)
		if err != nil {
			log.Println("hystrix error:", err.Error())
			time.Sleep(1 * time.Second)
		}
	}

	time.Sleep(100 * time.Second)
}
