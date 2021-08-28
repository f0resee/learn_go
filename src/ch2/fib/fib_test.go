package fib

import "testing"

func TestFibList(t *testing.T) {
	var a int = 1
	var b int = 1
	for i := 0; i < 10; i++ {
		t.Log(a)

		b = a + b
		a = b - a
	}
}
//multi assignment
func TestExchange(t *testing.T)  {
	var a int = 1
	var b int = 2
	a, b = b, a
	t.Log(a,b)
}

