package map_test

import "testing"

func TestMap(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	t.Log(len(m), m["one"])
	m1 := map[int]int{}
	t.Log(len(m1))
	m2 := make(map[int]int, 10)
	t.Log(len(m2))

	//judge if exist
	if v, ok := m2[3]; ok {
		t.Log(v)
	} else {
		t.Log("not exist")
	}

	//range
	for i,v:=range m{
		t.Log(i,v)
	}

}
