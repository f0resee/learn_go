package errors

import (
	//"errors"
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

// 1. WithStack
func Test_Errors_WithStack(t *testing.T) {
	err := errors.WithStack(fmt.Errorf("there is something wrong"))
	fmt.Printf("got error: %+v\n", err)
}

// 2. Wrap
func Test_Errors_Wrap(t *testing.T) {
	err0 := errors.New("error 0")
	fmt.Printf("error 0: %s\n", err0.Error())

	err1 := errors.Wrap(err0, "error 1")
	fmt.Printf("error 1: %s\n", err1.Error())

	err2 := errors.Wrap(err1, "error 2")
	fmt.Printf("error 2: %s\n", err2.Error())
}

// 3. Unwrap
func Test_Errors_Unwrap(t *testing.T) {
	err0 := errors.New("error 0")
	fmt.Printf("error 0: %s\n", err0.Error())

	err1 := fmt.Errorf("error 1: %w", err0)
	fmt.Printf("error 1: %s\n", err1.Error())

	err2 := errors.Unwrap(err1)
	fmt.Printf("error 2: %s\n", err2.Error())
}

// 4. Cause
func Test_Errors_Cause(t *testing.T) {
	err0 := errors.New("error 0")
	err1 := errors.Wrap(err0, "error 1")
	err2 := errors.Wrap(err1, "error 2")
	fmt.Printf("%s\n", err2.Error())
	fmt.Printf("%s\n", errors.Cause(err2).Error())
}

// 5. Is
func Test_Errors_Is(t *testing.T) {
	err0 := errors.New("error 0")
	err1 := fmt.Errorf("error 1: %w", err0)
	err2 := fmt.Errorf("error 2: %w", err1)

	fmt.Printf("%s\n", err2.Error())
	fmt.Printf("%v\n", errors.Is(err2, err1))
}

// 6. As
type TypicalErr struct {
	e string
}

func (t TypicalErr) Error() string {
	return t.e
}

// 输出：
// TypicalErr is on the chain of err2
// true
func Test_Errors_As(t *testing.T) {
	err := TypicalErr{"typical error"}
	err1 := fmt.Errorf("wrap err: %w", err)
	err2 := fmt.Errorf("wrap err1: %w", err1)
	var e TypicalErr
	if !errors.As(err2, &e) {
		panic("TypicalErr is not on the chain of err2")
	}
	println("TypicalErr is on the chain of err2")
	println(err == e)
}
