package strings_test

import "testing"

func TestStrings(t *testing.T) {
	var s string
	t.Log(s)
	s = "hello"
	t.Log(len(s))
	t.Log(s)
	s = "\xE4\xB8\xA5"
	t.Log(s)
	t.Log(len(s))

	c := []rune(s)
	t.Logf("unicode %x", c[0])
	t.Logf("utf8 %x", s)
}
func TestStringToRune(t *testing.T)  {
	s := "中华人民共和国"
	for _,c := range s{
		t.Logf("%[1]c %[1]d",c)
	}
}