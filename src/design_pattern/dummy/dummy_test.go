package dummy

import "testing"

func TestDummy(t *testing.T) {
	db := db{}
	a := db.GetCustomers("a")
	t.Log(a.GetName())
	b := db.GetCustomers("b")
	t.Log(b.GetName())
	dummy := db.GetCustomers("c")
	t.Log(dummy.GetName())
}
