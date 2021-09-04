package gomutex

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var m sync.Mutex
var set = make(map[int]int,0)

func PrintOnce(n int)  {
	m.Lock()
	defer m.Unlock()
	if _,ok:=set[n];ok{

	}else{
		fmt.Println(n)
		set[n] = 1
	}
}
func TestMutex(t *testing.T)  {
	for i:=0;i<10;i++{
		PrintOnce(100)
	}
	time.Sleep(time.Second)
}