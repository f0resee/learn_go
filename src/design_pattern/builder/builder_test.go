package builder

import (
	"encoding/json"
	"testing"
)

func TestBuilder(t *testing.T) {
	req := NewRequestBuilder().
		Method("GET").
		Header(map[string]string{
			"a": "a1",
			"b": "b1",
		}).Proto("http").
		Build()
	d, _ := json.Marshal(req)
	t.Log(string(d))
}
