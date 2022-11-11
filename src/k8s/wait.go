package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
stop channel
context
mock
http request
channel
*/

/*
核心：一个协程监听有没有收到signal，并通过关闭stop channel使其不阻塞来通知另外的协程
	另外一个协程读取stop channel来判断context是否应该结束
*/

var shutdownSignals = []os.Signal{syscall.SIGINT, syscall.SIGTERM}

var onlyOneSignalHandler = make(chan struct{})

func SetupSignalHandler() <-chan struct{} {
	return SetupSignalHandlerChan()
}

func SetupSignalHandlerChan() chan struct{} {
	close(onlyOneSignalHandler)

	stop := make(chan struct{})
	c := make(chan os.Signal, 2)
	signal.Notify(c, shutdownSignals...)
	go func() {
		s := <-c // 1. 在c收到signal之前，此处一直阻塞
		fmt.Printf("receive signal:%s, start to stop", s)
		close(stop)
		s = <-c // 收到一次signal后再次阻塞，实际上是给程序留下优雅退出的时间
		fmt.Printf("receive signal %s, exit", s)
		os.Exit(1)
	}()
	return stop
}

func ContextFromChannel(ch <-chan struct{}) (context.Context, context.CancelFunc) {
	return MergeChannelIntoContext(context.Background(), ch)
}

func MergeChannelIntoContext(ctx context.Context, ch <-chan struct{}) (context.Context, context.CancelFunc) {
	c, cancel := context.WithCancel(ctx)

	go func() {
		select {
		case d, ok := <-ch: // 2. 35行阻塞导致此处阻塞，而一旦收到了signal，channel被关闭，此处反而可以读取，并调用cancel
			fmt.Println("d = ", d, " ok =", ok)
			cancel()
		case <-c.Done():
		}
	}()
	return c, cancel
}

func main() {
	stopCh := SetupSignalHandler()
	ctx, cancel := ContextFromChannel(stopCh) // 后续使用的context应该从此处派生
	go func(c context.Context) {
		select {
		case <-c.Done():
			fmt.Println("ctx done. routine exit.")
			return
		}
	}(ctx)
	go func() {
		time.Sleep(10 * time.Second)
		fmt.Println("10s passed. program want to exit.")
		cancel() // 等效于发送signal，可以在遇到不可恢复或者缺少必要条件时主动退出
	}()
	time.Sleep(20 * time.Second)
}
