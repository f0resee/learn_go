package gochannel

import (
	"fmt"
	"testing"
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
