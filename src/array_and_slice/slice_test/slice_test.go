package slice_test

import "testing"

/*
golang slice 扩容策略：
	1. 如果申请的容量大于两倍旧的容量，直接扩容至申请容量
	2. 否则，如果旧切片的长度小于1024，直接扩容至原来的两倍
	3. 否则，如果旧切片长度大于等于1024，最终容量从原来的容量开始，每次增加原来的1/4，直到大于申请的容量
	4. 如果最终容量计算值溢出，最终容量就是新申请容量
*/

func TestSlice(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))

	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))

	s2 = append(s2, 1)
	t.Log(len(s2), cap(s2))
	t.Log(s2[3])

}

func TestSliceGrow(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceSharedMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	q2 := year[3:6]
	t.Log(q2, len(q2), cap(q2))

	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))

	summer[0] = "Unknow"
	t.Log(q2)

}
func TestSliceCompare(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{1, 2, 3, 4}
	if a == nil || b == nil {
		t.Log("nil")
	} else {
		t.Log("not nil")
	}
}

func TestSliceModify(t *testing.T) {
	t.Run("just modify", func(t *testing.T) {
		modifySlice := func(s []int) {
			s[0] = 1024
		}
		var s []int
		for i := 0; i < 3; i++ {
			s = append(s, i)
		}
		modifySlice(s)
		t.Log(s)
	})

	t.Run("append and modify", func(t *testing.T) {
		modifySlice := func(s []int) {
			s = append(s, 2048)
			s[0] = 1024
		}
		var s []int
		for i := 0; i < 3; i++ {
			s = append(s, i)
		}
		modifySlice(s)
		t.Log(s)
	})

	t.Run("append scale and modify", func(t *testing.T) {
		modifySlice := func(s []int) {
			s = append(s, 2048)
			s = append(s, 4096) //发生扩容
			s[0] = 1024         //修改的是扩容后的slice
			//t.Log(cap(s))
		}
		var s []int
		for i := 0; i < 3; i++ {
			s = append(s, i)
		}
		//t.Log(cap(s))
		modifySlice(s)
		t.Log(s)
	})
}

func BenchmarkName(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n++
	}
}
