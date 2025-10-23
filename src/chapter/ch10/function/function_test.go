package function_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int,int) {
	return rand.Intn(10),rand.Intn(20)
}
func TestFunction(t *testing.T)  {
	t.Log(returnMultiValues())
}

func timeSpent(inner func(op int)int) func(op int) int  {
	return func(n int) int {
		start := time.Now()
		ret:=inner(n)
		fmt.Println("Time spent:",time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(op int)int  {
	time.Sleep(time.Second*1)
	return op
}

func TestTimeSpent(t *testing.T)  {
	a,_:=returnMultiValues()
	t.Log(a)
	tsf :=timeSpent(slowFunc)
	t.Log(tsf(10))
}

func Sum(ops... int)int  {
	ret := 0
	for _,op :=range ops{
		ret = ret + op
	}
	return ret
}
func TestVarParam(t *testing.T)  {
	t.Log(Sum(1,2,3,4))
	t.Log(Sum(1,2,3,4,5))
}

func Clear()  {
	fmt.Println("Clear resource")
}
func TestDefer(t *testing.T)  {
	defer func() {
		t.Log("clear resources")
	}()
	//defer Clear()
	t.Log("start")
	panic("Fatal error")
}