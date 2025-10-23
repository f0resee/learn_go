package interface_test

import "testing"

type Code string
type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {

}

func (g *GoProgrammer)WriteHelloWorld() string  {
	return "fmt.Println(\"Hello, World\")"
}

func TestProgrammer(t *testing.T)  {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}