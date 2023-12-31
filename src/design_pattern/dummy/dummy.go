package dummy

type AbstractCustomer interface {
	GetName() string
}

type RealCustomer struct {
	name string
}

func (r *RealCustomer) GetName() string {
	return r.name
}

type Dummy struct {
}

func (r *Dummy) GetName() string {
	return ""
}

type db struct {
}

func (d *db) GetCustomers(name string) AbstractCustomer {
	if name == "a" || name == "b" {
		return &RealCustomer{
			name: name,
		}
	}
	return &Dummy{}
}
