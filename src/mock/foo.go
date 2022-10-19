package mock

import "fmt"

type Foo interface {
	Bar(x int) int
}

func SUT(f Foo) {
	fmt.Println("number is ", f.Bar(99))
}
