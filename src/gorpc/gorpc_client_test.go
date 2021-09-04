package gorpc

import (
	"net/rpc"
	"testing"
)

func TestRpcClient(t *testing.T) {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		t.Log("dialing:", err)
	}
	args := &Args{7, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		t.Log("arith error:", err)
	}
	t.Logf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	quotient := new(Quotient)
	divCall := client.Go("Arith.Divide", args, quotient, nil)
	replyCall := <-divCall.Done
	if replyCall.Error != nil {
		t.Log("arith error:", replyCall.Error)
	}
	t.Logf("Arith: %d/%d=%d...%d", args.A, args.B, quotient.Quo, quotient.Rem)
}
