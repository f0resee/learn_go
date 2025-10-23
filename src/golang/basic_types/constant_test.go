package constant_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const(
	Readable = 1<<iota
	Writable
	Executable
)
func TestConstantTry(t *testing.T)  {
	t.Log(Monday,Tuesday)

}
func TestConstantTry1(t *testing.T) {
	t.Log(Readable,Writable,Executable)
}
