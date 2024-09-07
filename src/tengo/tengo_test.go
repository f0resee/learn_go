package tengo

import (
	"fmt"
	"testing"

	"github.com/d5/tengo/v2"
)

type Parameters interface {
	GetParameters() map[string]interface{}
}

func NewParameters(params map[string]interface{}) Parameters {
	p := make(map[string]interface{})
	for key, value := range params {
		p[key] = value
	}
	return &parameters{
		parameters: p,
	}
}

type parameters struct {
	parameters map[string]interface{}
}

func (p *parameters) GetParameters() map[string]interface{} {
	return p.parameters
}

type Result interface {
	GetBool(name string) (bool, error)
	GetString(name string) (string, error)
	GetInt(name string) (int64, error)
}

func NewResult(compiled *tengo.Compiled) Result {
	return &result{
		compiled: compiled,
	}
}

type result struct {
	compiled *tengo.Compiled
}

func (r *result) GetBool(name string) (bool, error) {
	if !r.compiled.IsDefined(name) {
		return false, fmt.Errorf("%s is undefined in compiled script", name)
	}
	value := r.compiled.Get(name)
	ret, ok := value.Value().(bool)
	if !ok {
		return false, fmt.Errorf("value %v is not instance of bool", value.Value())
	}
	return ret, nil
}

func (r *result) GetString(name string) (string, error) {
	if !r.compiled.IsDefined(name) {
		return "", fmt.Errorf("%s is undefined in compiled script", name)
	}
	value := r.compiled.Get(name)
	ret, ok := value.Value().(string)
	if !ok {
		return "", fmt.Errorf("value %v is not instance of string", value.Value())
	}
	return ret, nil
}

func (r *result) GetInt(name string) (int64, error) {
	if !r.compiled.IsDefined(name) {
		return 0, fmt.Errorf("%s is undefined in compiled script", name)
	}
	value := r.compiled.Get(name)
	ret, ok := value.Value().(int64)
	if !ok {
		return 0, fmt.Errorf("value %v is not instance of int", value.Value())
	}
	return ret, nil
}

type TengoProgram interface {
	Run(params Parameters) (Result, error)
}

func NewTengoProgram(script string) TengoProgram {
	return &tengoProgram{
		script: tengo.NewScript([]byte(script)),
	}
}

type tengoProgram struct {
	script *tengo.Script
}

func (t *tengoProgram) Run(params Parameters) (Result, error) {
	for key, value := range params.GetParameters() {
		err := t.script.Add(key, value)
		if err != nil {
			return nil, err
		}
	}
	compiled, err := t.script.Run()
	if err != nil {
		return nil, err
	}

	return NewResult(compiled), nil
}

func TestTengo(t *testing.T) {
	script := `each := func(seq, fn) {
    for x in seq { fn(x) }
}

sum := 0
mul := 1
each([a, b, c, d], func(x) {
    sum += x
    mul *= x
})`

	params := map[string]interface{}{
		"a": 1,
		"b": 9,
		"c": 8,
		"d": 4,
	}

	program := NewTengoProgram(script)
	result, err := program.Run(NewParameters(params))
	if err != nil {
		t.Error(err)
	}

	sum, err := result.GetInt("sum")
	if err != nil {
		t.Error(err)
	}
	t.Log("sum = ", sum)
	mul, err := result.GetInt("mul")
	if err != nil {
		t.Error(err)
	}
	t.Log("mul = ", mul)

	script = `
	ret := "return"
`
	pro := NewTengoProgram(script)
	re, err := pro.Run(NewParameters(map[string]interface{}{}))
	if err != nil {
		t.Error(err)
	}
	t.Log(re.GetString("ret"))

}
