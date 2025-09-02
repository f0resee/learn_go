package gochannel

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	c := make(chan int)
	defer close(c)
	go func() { c <- 3 + 4 }()
	i, ok := <-c
	fmt.Println(i, ok)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}
func TestGoChannel(t *testing.T) {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x+y)
}

func TestSingleDirectionChannel(t *testing.T) {
	t.Run("no buffer channel", func(t *testing.T) {
		c := make(chan int) // 无缓存的channel，写入是阻塞的
		start := time.Now()
		go func(ch chan<- int) { // 只写的channel，只读的为 <-chan int
			ch <- 1
			t.Logf("go routine exit, time elapsed:%v", time.Since(start))
		}(c)
		time.Sleep(1 * time.Second)
		<-c
	})
	t.Run("buffer channel", func(t *testing.T) {
		c := make(chan int, 1) //有缓存的channel，写入不阻塞
		start := time.Now()

		go func(ch chan<- int) {
			ch <- 1
			t.Logf("go routine exit, time elapsed:%v", time.Since(start))
		}(c)
		time.Sleep(1 * time.Second)
		<-c
	})
	t.Run("read only channel", func(t *testing.T) {
		c := make(chan int, 1) //有缓存的channel，写入不阻塞
		start := time.Now()

		go func(ch <-chan int) {
			d := <-ch
			t.Logf("go routine exit, time elapsed: %v, data = %v", time.Since(start), d)
		}(c)
		c <- 1
		time.Sleep(1 * time.Second)
	})
}

func TestCloseChannel(t *testing.T) {
	ch := make(chan struct{})
	go func(c chan struct{}) {
		select {
		case <-c:
			t.Log("receive")
		}
	}(ch)
	time.Sleep(1 * time.Second)
	close(ch)
	time.Sleep(1 * time.Second)
}
