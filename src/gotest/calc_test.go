package gotest

import "testing"

func TestAdd(t *testing.T) {
	if ans := Add(1, 2); ans != 3 {
		t.Errorf("1 + 2 expected to be 3, but %d got", ans)
	}

	if ans := Add(-10, -20); ans != -30 {
		t.Errorf("-10 + -20 expected to be -30, but %d got", ans)
	}
}

func TestMul(t *testing.T) {
	t.Run("pos", func(t *testing.T) { // 子测试
		if Mul(2, 3) != 6 {
			t.Fatal("fail")
		}
	})

	t.Run("neg", func(t *testing.T) {
		if Mul(2, -3) != -6 {
			t.Fatal("fail")
		}
	})
}

func TestMul2(t *testing.T) {
	cases := []struct {
		name           string
		a, b, expected int
	}{
		{"pos", 2, 3, 6},
		{"neg", 2, -3, -6},
		{"zero", 2, 0, 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if ans := Mul(c.a, c.b); ans != c.expected {
				t.Fatalf("%d * %d expected %d, but got %d", c.a, c.b, c.expected, ans)
			}
		})
	}
}
