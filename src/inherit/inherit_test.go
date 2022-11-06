package inherit

import (
	"fmt"
	"testing"
)

type Number interface {
	Print()
}

type Int struct {
	num int
}

func (i *Int) Print() {
	fmt.Println(i.num)
}

type Pair struct {
	Number
	f float32
}

func TestInherit(t *testing.T) {
	p := Pair{
		Number: &Int{
			num: 1,
		},
	}
	p.Print()
}
