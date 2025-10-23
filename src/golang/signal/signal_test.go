package signal

import (
	"sync"
	"testing"
)

var once sync.Once

var onlyOneSignalHanler = make(chan struct{})

func TestSignal(t *testing.T) {
	close(onlyOneSignalHanler)
}
