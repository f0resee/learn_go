package basictypes

import (
	"reflect"
	"testing"
)

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
	for i, v := range m {
		t.Log(i, v)
	}

}

// 比较两个map是否相等
func cmpMap(map1, map2 map[string]string) bool {
	for k, v := range map1 {
		if v2, exist := map2[k]; !exist || v != v2 {
			return false
		}
	}
	for k, v := range map2 {
		if v1, exist := map1[k]; !exist || v != v1 {
			return false
		}
	}
	return true
}

// map比较相等
func TestMapCompare(t *testing.T) {
	tests := []struct {
		name   string
		mp1    map[string]string
		mp2    map[string]string
		result bool
	}{
		{
			name: "equal",
			mp1: map[string]string{
				"abc": "def",
				"bcd": "efg",
				"cde": "fgh",
			},
			mp2: map[string]string{
				"abc": "def",
				"bcd": "efg",
				"cde": "fgh",
			},
			result: true,
		},
		{
			name: "not equal",
			mp1: map[string]string{
				"abc": "def",
				"bcd": "efg",
				"cde": "fgh",
			},
			mp2: map[string]string{
				"abc": "abc",
				"bcd": "bcd",
				"cde": "cde",
			},
			result: false,
		},
	}

	//通过key-value逐项比较
	for _, tt := range tests {
		t.Run(tt.name+" by key-value", func(t *testing.T) {
			if r := cmpMap(tt.mp1, tt.mp2); r != tt.result {
				t.Errorf("result is %v, want %v", r, tt.result)
			}
		})
	}

	//使用DeepEqual
	for _, tt := range tests {
		t.Run(tt.name+" by DeepEqual", func(t *testing.T) {
			if r := reflect.DeepEqual(tt.mp1, tt.mp2); r != tt.result {
				t.Errorf("result is %v, want %v", r, tt.result)
			}
		})
	}
}
