package filter

import "testing"

func TestFilterChain(t *testing.T) {
	chain := NewFilterChain()
	chain.AddFilter(&authenticationFilter{})
	chain.AddFilter(&debugFilter{})
	chain.Execute("test")
}
