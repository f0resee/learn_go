package filter

import "fmt"

type Filter interface {
	Execute(req string)
}

type authenticationFilter struct {
}

func (a *authenticationFilter) Execute(req string) {
	fmt.Printf("authentication req: %s\n", req)
}

type debugFilter struct {
}

func (d *debugFilter) Execute(req string) {
	fmt.Printf("debug req: %s\n", req)
}

func NewFilterChain() FilterChain {
	return FilterChain{
		filters: make([]Filter, 0),
	}
}

type FilterChain struct {
	filters []Filter
}

func (c *FilterChain) AddFilter(filter Filter) {
	c.filters = append(c.filters, filter)
}

func (c *FilterChain) Execute(req string) {
	for _, filter := range c.filters {
		filter.Execute(req)
	}
}
