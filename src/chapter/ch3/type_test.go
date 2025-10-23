package type_test

import "testing"

/*
bool
string
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr
byte //uint8
rune //int32
float32 float64
complex64 complex128

no implicit conversion, even for alias

*/
type myint int64

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	//b = a
	b = int64(a)
	t.Log(a, b)
	var c myint = 0
	c = myint(b)
	t.Log(b, c)
}

//no pointer calc
func TestPointer(t *testing.T)  {
	a := 1
	aPtr := &a
	t.Log(a,aPtr)
	t.Logf("%T %T",a,aPtr)
}

func TestString(t *testing.T)  {
	var s string
	t.Log(s)
}