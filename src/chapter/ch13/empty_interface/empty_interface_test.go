package empty_interface__test

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	/*
		if i, ok := p.(int); ok {
			fmt.Println("int ", i)
			return
		}
		if s, ok := p.(string); ok {
			fmt.Println("string ", s)
			return
		}
		fmt.Println("Unknown type")
	*/
	switch v:=p.(type) {
	case int:
		fmt.Println("int ", v)
		return
	case string:
		fmt.Println("string ", v)
		return
	default:
		fmt.Println("Unknown ")
		return
	}
}
func TestEmptyInterface(t *testing.T) {
	DoSomething(10)
	DoSomething("hello")
}
