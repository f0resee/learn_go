package gosync

import (
	"fmt"
	"sync"
	"testing"
)

func TestGoOnce(t *testing.T)  {
	var once sync.Once
	onceBody := func() {
		fmt.Println("only once")
	}
	done := make(chan bool)
	for i:=0;i<10;i++ {
		go func() {
			once.Do(onceBody)
			done<-true
		}()
	}
	for i:=0;i<10;i++{
		<-done
	}
}
