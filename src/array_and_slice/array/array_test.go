package array_test

import "testing"

func TestArrayInit(t *testing.T)  {
	var a [3]int
	t.Log(a[0],a[1],a[2])

	arr1 := [4]int{1,2,3,4}
	t.Log(arr1[0],arr1[1],arr1[2],arr1[3])

	arr2 := [...]int{1,2,3,5}
	t.Log(arr2[0],arr2[1],arr2[2],arr2[3])
	/*
	for idx,e:=range arr2{
			t.Log(idx,e)
		}
	 */
	for _,e:=range arr2{
		t.Log(e)
	}

	//a[index1:index2]
	//[index1,index2)
	arr3 := arr1[1:]
	for _,e:=range arr3{
		t.Log(e)
	}

	var arr4 = make([]int,10)
	t.Log(arr4[3])
}