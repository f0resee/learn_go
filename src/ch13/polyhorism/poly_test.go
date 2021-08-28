package polyhorism__test

import (
	"fmt"
	"testing"
)

type Code string
type Programmer interface {
	WriteHelloWorld() string
}

func WriteFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello, World\")"
}

type JavaProgrammer struct {
}

func (g *JavaProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello, World\")"
}
func TestProgrammer(t *testing.T) {
	//var p Programmer
	p := new(GoProgrammer)
	q := new(JavaProgrammer)
	WriteFirstProgram(p)
	WriteFirstProgram(q)
}
