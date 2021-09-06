package gosync

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestGoSync(t *testing.T) {
	c := sync.NewCond(&sync.Mutex{})
	var ready int

	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)
			c.L.Lock()
			ready++
			c.L.Unlock()
			t.Logf("%d start", i)
			//c.Signal()
			c.Broadcast()

		}(i)
	}

	c.L.Lock()
	for ready != 10 {
		//Wait atomically unlocks c.L and suspends execution of the calling goroutine.
		//After later resuming execution, Wait locks c.L before returning.
		//Unlike in other systems, Wait cannot return unless awoken by Broadcast or Signal.
		c.Wait()
		t.Logf("await one time")
	}
	c.L.Unlock()
	t.Logf("all started")
}
