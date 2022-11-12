package wait

import (
	"context"
	"testing"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
)

func TestPollUntil(t *testing.T) {
	ctx := context.Background()
	toCtx, _ := context.WithTimeout(ctx, 10*time.Second)
	t0 := time.Now()
	go wait.PollUntil(3*time.Second, func() (done bool, err error) { // 经过一次interval之后才会执行
		t.Logf("time passed:%v", time.Since(t0))
		return false, nil
	}, toCtx.Done())
	time.Sleep(15 * time.Second)
}

func TestUntil(t *testing.T) {
	ctx := context.Background()
	ttCtx, _ := context.WithTimeout(ctx, 10*time.Second)
	t0 := time.Now()
	go wait.Until(func() {
		t.Logf("time passed:%v", time.Since(t0))
		return
	}, 1*time.Second, ttCtx.Done())
	time.Sleep(15 * time.Second)
}

func TestPoll(t *testing.T) {
	t0 := time.Now()
	wait.Poll(1*time.Second, 10*time.Second, func() (done bool, err error) {
		t.Logf("time passed:%v", time.Since(t0))
		return false, nil
	})
}

func TestWaitFor(t *testing.T) {
	//wait.WaitFor()
}
