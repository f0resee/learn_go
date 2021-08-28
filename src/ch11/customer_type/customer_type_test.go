package customer_type

import (
	"fmt"
	"testing"
	"time"
)

type IntConv func(op int) int

func timeSpent(inner IntConv) IntConv {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("Time spent:", time.Since(start).Seconds())
		return ret
	}
}
func slowFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestTimeSpent(t *testing.T) {
	tsf := timeSpent(slowFunc)
	t.Log(tsf(10))
}
