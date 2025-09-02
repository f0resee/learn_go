package singnal

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestSignal(t *testing.T) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("start")
	s := <-c
	fmt.Println("finish: ", s.String())
}
