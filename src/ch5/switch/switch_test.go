package switch_test

import "testing"

func TestSwitch(t *testing.T) {
	var a int = 1
	switch a {
	case 0, 1:
		t.Log("a<2")
	case 2:
		t.Log("a==2")
	default:
		t.Log("a>2")
	}
}
func TestSwitch1(t *testing.T) {
	var a int = 1
	switch {
	case a%2 == 0:
		t.Log("Even")
	default:
		t.Log("Odd")
	}
}
