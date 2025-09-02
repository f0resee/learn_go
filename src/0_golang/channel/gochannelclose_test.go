package gochannel

import (
	"fmt"
	"sync"
	"testing"
)

func dataConsumer(ch chan int,wg *sync.WaitGroup)  {
	go func() {
		for  {
			if data,ok:=<-ch;ok{
				fmt.Println(data)
			}else{
				fmt.Println(<-ch) //default
				fmt.Println("channel closed")
				break
			}
		}
		wg.Done()
	}()
}

func dataProducer(ch chan int,wg *sync.WaitGroup)  {
	go func() {
		for i:=0;i<10;i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()
}

func TestChannelClose(t *testing.T)  {
	var wg sync.WaitGroup
	ch := make(chan int,1)
	wg.Add(1)
	dataProducer(ch,&wg)
	wg.Add(1)
	dataConsumer(ch,&wg)
	wg.Add(1)
	dataConsumer(ch,&wg)
	wg.Wait()
}