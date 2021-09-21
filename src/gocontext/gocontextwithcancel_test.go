package gocontext

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"
)

var logg *log.Logger

func someHandler() {
	ctx, cancel := context.WithCancel(context.Background())
	go doStuff(ctx)

	//10秒后取消doStuff
	time.Sleep(10 * time.Second)
	cancel()
	time.Sleep(1*time.Second) //等待协程退出，不等待看不到后续输出
}

//每1秒work一下，同时会判断ctx是否被取消了，如果是就退出
func doStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)

		select {
		case <-ctx.Done():  //写法非常奇怪，10s后cancel实际上是永远无法满足此条件的
			logg.Printf("done")
			return
		default:
			logg.Printf("work")
		}

		logg.Printf("WORK")
	}
}

func TestWithCancel(t *testing.T) {
	logg = log.New(os.Stdout, "", log.Ltime)
	someHandler()
	logg.Printf("down")
}

func TestCotextWithCancel(t *testing.T) {
	// gen generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the context once
	// they are done consuming generated integers not to leak
	// the internal goroutine started by gen.
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					t.Log("EXIT")
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

func TestContextDone(t *testing.T)  {
	ctx, cancel := context.WithCancel(context.Background())
	i := 0

	// monitor

	go func() {
		for range time.Tick(time.Second) {
			select {
			case <-ctx.Done():
				//fmt.Println("exiting")
				i = i + 1
				fmt.Println(i)
				return
			default:
				fmt.Println("monitor working")
			}
		}
	}()

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(time.Second)
	//fmt.Println("i=",i)
}