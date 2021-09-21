package gocontext

import (
	"context"
	"fmt"
	"testing"
	"time"
)

const shortDuration = 1 * time.Millisecond

func TestWithDeadLine(t *testing.T) {
	d := time.Now().Add(shortDuration)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancellation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func TestWithTimeout(t *testing.T)  {
	ctx, cancel := context.WithTimeout(context.Background(),time.Second)
	defer func() {
		//t.Log("trying to cancel")
		cancel()   //一旦执行cancel()，子任务会立即退出
	}()
	go func(ctx context.Context) {
		for  {
			time.Sleep(time.Millisecond*95)
			fmt.Println("working")
		}
	}(ctx)

	select {
	case <-ctx.Done():
		t.Log(ctx.Err())
	case <-time.After(time.Second*5):
		t.Log("working")
	}
}