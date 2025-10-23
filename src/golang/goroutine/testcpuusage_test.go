package goroutine

import (
	"math"
	"testing"
)

func isPrime(i int64) bool {
	count := 0
	var j int64 = 2
	for ; j < i; j++ {
		if i%j == 0 {
			count++
		}
	}
	if count == 0 {
		return true
	}
	return false
}

func TestCpuUsage(t *testing.T) {
	for i := 0; i < 10000; i++ {
		go func() {
			var j int64 = 2
			for ; j < math.MaxInt64; j++ {
				if isPrime(j) {
					println(j)
				}
			}
		}()
	}
	t.Log("finished")
}
