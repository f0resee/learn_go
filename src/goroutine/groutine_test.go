package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

func gPrint(id int,wg *sync.WaitGroup) {
	defer wg.Done()//计数器减1
	fmt.Printf("routine id is %d\n", id)
}
// 协程使用waitgroup
func TestRoutine(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1) //计数器加1
		go gPrint(i,&wg)
	}
	wg.Wait() //等待计数器归零
}
