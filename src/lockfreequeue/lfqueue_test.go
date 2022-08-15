package lockfreequeue

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestLFQueue(t *testing.T) {
	lfQueue := NewLKQueue()
	tStart := time.Now()
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 10000000; j++ {
				lfQueue.Enqueue(rand.Int())
			}
			wg.Done()
		}()
	}
	wg.Wait()
	t.Log("time elapsed ", time.Since(tStart).Milliseconds())
}
func TestLockQueue(t *testing.T) {
	lfQueue := MakeLockQueue()
	tStart := time.Now()
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 10000000; j++ {
				lfQueue.Enqueue(rand.Int())
			}
			wg.Done()
		}()
	}
	wg.Wait()
	t.Log("time elapsed ", time.Since(tStart).Milliseconds())
}
