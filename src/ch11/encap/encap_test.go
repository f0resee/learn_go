package encap_test

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func (e *Employee) String() string {
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}
//func (e *Employee) String1() string {
func (e *Employee) String1() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

func TestEmployee(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	t.Log(e.String())
	fmt.Printf("Address is %x\n",unsafe.Pointer(&e.Name))
	t.Log(e.String1())
}
