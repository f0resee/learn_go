package goerror

import (
	"errors"
	"fmt"
	"testing"
)

func TestStdErrors(t *testing.T) {
	err1 := errors.New("error1")
	err2 := fmt.Errorf("error2: [%w]", err1)
	fmt.Println(err2)
	fmt.Println(errors.Unwrap(err2))
	// Output
	// error2: [error1]
	// error1
}
