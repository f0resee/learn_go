package gocontext

import (
	"context"
	"fmt"
	"testing"
)

func TestContextWithValue(t *testing.T) {
	type favContextKey string

	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	f(ctx, k)
	f(ctx, favContextKey("color"))


	k2 := favContextKey("language")
	ctx2 := context.WithValue(context.Background(),k2,"LANG")
	f(ctx2,k2)

}
